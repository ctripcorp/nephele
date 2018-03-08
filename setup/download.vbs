WSH.Echo "downloading,please wait!"
Set xPost = CreateObject("Microsoft.XMLHTTP")  
xPost.Open "GET","http://file.c-ctrip.com/files/3/car22/zip/7a7/cf5/e08/a24538a1ed74420f95745097fa987fd8.zip",0   
xPost.Send()  
Set sGet = CreateObject("ADODB.Stream")  
sGet.Mode = 3  
sGet.Type = 1  
sGet.Open()  
sGet.Write(xPost.responseBody)  
sGet.SaveToFile "gm.zip",2  
WSH.Echo "download end"
wscript.sleep 1000

Class ZipCompressor
 
    Private objFileSystemObject
    Private objShellApplication
    Private objWScriptShell
    Private objScriptingDictionary
    Private objWMIService
    Private COPY_OPTIONS
 
    Private Sub Class_Initialize()
        Set objFileSystemObject = WSH.CreateObject("Scripting.FileSystemObject")
        Set objShellApplication = WSH.CreateObject("Shell.Application")
        Set objWScriptShell     = WSH.CreateObject("WScript.Shell")
        Set objScriptingDictionary = WSH.CreateObject("Scripting.Dictionary")
        Dim strComputer
        strComputer = "."
        Set objWMIService = GetObject("winmgmts:\\" & strComputer & "\root\cimv2")
 
        ' COPY_OPTIONS
        '    4   Do not display a progress dialog box.
        '   16   Respond with "Yes to All" for 
        '         any dialog box that is displayed.
        '  512   Do not confirm the creation of a new 
        '         directory if the operation requires one to be created.
        ' 1024   Do not display a user interface if an error occurs.
 
        COPY_OPTIONS =  4 + 16 + 512 + 1024
    End Sub
 
    Private Sub Class_Terminate()
        Set objWMIService = Nothing
        objScriptingDictionary.RemoveAll
        Set objScriptingDictionary = Nothing
        Set objWScriptShell     = Nothing
        Set objShellApplication = Nothing
        Set objFileSystemObject = Nothing
    End Sub
 
 
    Private Sub makeEmptyZipFile(pathToZipFile)
        Dim file
        Set file = objFileSystemObject.CreateTextFile(pathToZipFile)
        file.Write Chr(80) & Chr(75) & Chr(5) & Chr(6) & String(18, 0)
        file.Close
    End Sub
 
    Private Function pathToAbsolute(fileName)
        Dim i, file, files
        files = Split(fileName, ";")
        ReDim tmpFiles(UBound(files))
 
        i = 0
        For Each file in files
            If file<>"" Then
                file = objWScriptShell.ExpandEnvironmentStrings(file)
                file = objFileSystemObject.GetAbsolutePathName(file)
 
                tmpFiles(i) = file
                i = i+1
            End If
        Next
        If i-1 > 0 And i-1 < UBound(files) Then ReDim Preserve tmpFiles(i-1)
        pathToAbsolute = Join(tmpFiles, ";")
        Erase tmpFiles
    End Function
 
    Private Function pathCombine(fileName, nextFileName)
        Dim files, lastIndex
        files = Split(fileName, "\")
        lastIndex = UBound(files)
 
        If files(lastIndex)<>"" Then
            lastIndex = lastIndex + 1
            ReDim Preserve files(lastIndex)
        End If
 
        files(lastIndex) = nextFileName
 
        pathCombine = Join(files, "\")
        Erase files
    End Function
 
    Private Function pathSplit(fileName)
        Dim fileSplitted(2)
        fileSplitted(0) = objFileSystemObject.GetDriveName(fileName)
        fileSplitted(2) = objFileSystemObject.GetFileName(fileName)
        fileSplitted(1) = Mid(fileName, Len(fileSplitted(0))+1, _
            Len(fileName) - Len(fileSplitted(0)) - Len(fileSplitted(2)))
 
        pathSplit = fileSplitted
    End Function
 
    Private Function pathSplitForQuery(fileName)
        Dim fileSplitted
        fileSplitted = pathSplit(fileName)
        fileSplitted(1) = Replace(fileSplitted(1), "\", "\\")
        If Right(fileSplitted(1), 2) <> "\\" Then
            fileSplitted(1) = fileSplitted(1) & "\\"
        End If
        ' http://msdn.microsoft.com/en-us/library/windows/desktop/aa392263(v=vs.85).aspx
        fileSplitted(2) = Replace(fileSplitted(2), "_", "[_]")
        fileSplitted(2) = Replace(fileSplitted(2), "*", "%")
        fileSplitted(2) = Replace(fileSplitted(2), "?", "_")
        pathSplitForQuery = fileSplitted
    End Function
 
    Private Function buildQuerySQL(fileName)
        Dim fileSplitted, file, ext
        fileSplitted = pathSplitForQuery(fileName)
 
        Dim lastDotIndex
 
        file = "%" : ext  = "%"
        If fileSplitted(2)<>"" Then
            lastDotIndex = InStrRev(fileSplitted(2), ".")
            file = fileSplitted(2)
        End If
 
        If lastDotIndex>0 Then
            ext = Mid(fileSplitted(2), _
                lastDotIndex+1, Len(fileSplitted(2)) - lastDotIndex)
            file = Left(fileSplitted(2), Len(fileSplitted(2)) - Len(ext) - 1)
        End If
 
        ' http://msdn.microsoft.com/en-us/library/windows/desktop/aa387236(v=vs.85).aspx
        buildQuerySQL = "SELECT * FROM CIM_DataFile" & _
                        " WHERE Drive='" & fileSplitted(0) & "' AND" & _
                        " (FileName LIKE '" & file & "') AND" & _
                        " (Extension LIKE '" & ext & "') AND" & _
                        " (Path='" & fileSplitted(1) &"')"
    End Function
 
    Private Function deleteFile(fileName)
        deleteFile = False
        If objFileSystemObject.FileExists(fileName) Then
            objFileSystemObject.DeleteFile fileName
            deleteFile = True
        End If
    End Function
 
    Private Sub compress_(ByVal fileName, ByRef zipFile)
        Dim objFile, srcFile, srcFiles
        srcFiles = Split(fileName, ";")
 
        Dim colFiles
 
        ' http://msdn.microsoft.com/en-us/library/bb787866(VS.85).aspx
        For Each srcFile In srcFiles
            If objFileSystemObject.FolderExists(srcFile) Then
                Set objFile = objShellApplication.NameSpace(srcFile)
                If Not (objFile Is Nothing) Then
                    zipFile.CopyHere objFile.Items, COPY_OPTIONS
                    Do Until objFile.Items.Count <= zipFile.Items.Count
                        WScript.Sleep(200)
                    Loop
                End If
                Set objFile = Nothing
            ElseIf objFileSystemObject.FileExists(srcFile) Then
                zipFile.CopyHere srcFile, COPY_OPTIONS
                WScript.Sleep(200)
            Else
                Set colFiles = objWMIService.ExecQuery(buildQuerySQL(srcFile))
                For Each objFile in colFiles
                    srcFile = objFile.Name
                    zipFile.CopyHere srcFile, COPY_OPTIONS
                    WScript.Sleep(200)
                Next
                Set colFiles = Nothing
            End If
        Next
    End Sub
 
    Public Sub add(fileName)
        objScriptingDictionary.Add pathToAbsolute(fileName), ""
    End Sub
 
    ' Private Function makeTempDir()
    '    Dim tmpFolder, tmpName
    '    tmpFolder = objFileSystemObject.GetSpecialFolder(2)
    '    tmpName   = objFileSystemObject.GetTempName()
    '    makeTempDir = pathCombine(tmpFolder, tmpName)
    '    objFileSystemObject.CreateFolder makeTempDir
    ' End Function
 
    Public Function compress(srcFileName, desFileName)
        Dim srcAbsFileName, desAbsFileName
 
        srcAbsFileName = ""
        If srcFileName<>"" Then
            srcAbsFileName = pathToAbsolute(srcFileName)
        End If
 
        desAbsFileName = pathToAbsolute(desFileName)
 
        If objFileSystemObject.FolderExists(desAbsFileName) Then
            compress = -1
            Exit Function
        End If
 
        ' That zip file already exists - deleting it.
        deleteFile desAbsFileName
 
        makeEmptyZipFile desAbsFileName
 
        Dim zipFile
        Set zipFile = objShellApplication.NameSpace(desAbsFileName)
 
        If srcAbsFileName<>"" Then
            compress_ srcAbsFileName, zipFile
        End If
        compress = zipFile.Items.Count
 
        Dim objKeys, i
        objKeys = objScriptingDictionary.Keys
        For i = 0 To objScriptingDictionary.Count -1
            compress_ objKeys(i), zipFile
        Next
 
        compress = compress + i
 
        Set zipFile = Nothing
    End Function
 
    Public Function decompress(srcFileName, desFileName)
        Dim srcAbsFileName, desAbsFileName
        srcAbsFileName = pathToAbsolute(srcFileName)
        desAbsFileName = pathToAbsolute(desFileName)
 
        If Not objFileSystemObject.FileExists(srcAbsFileName) Then
            decompress = -1
            Exit Function
        End If
 
        If Not objFileSystemObject.FolderExists(desAbsFileName) Then
            decompress = -1
            Exit Function
        End If
 
        Dim zipFile, objFile
        Set zipFile = objShellApplication.NameSpace(srcAbsFileName)
            Set objFile = objShellApplication.NameSpace(desAbsFileName)
                objFile.CopyHere zipFile.Items, COPY_OPTIONS
                Do Until zipFile.Items.Count <= objFile.Items.Count
                    WScript.Sleep(200)
                Loop
 
                decompress = objFile.Items.Count
            Set objFile = Nothing
        Set zipFile = Nothing
    End Function
End Class

Dim zip
Set zip = New ZipCompressor
    ' 需要在D盘建立文件夹extract
    zip.decompress "gm.zip", "gm"
Set zip = Nothing

