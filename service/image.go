package service

// Group image handling all together
type ImageService struct {
	internal *Service
	routeMap map[string]string
}

// Modify get image path.
func (s *ImageService) Get(relativePath string) {
	s.routeMap["GET"] = relativePath
}

// Modify upload image path.
func (s *ImageService) Upload(relativePath string) {
	s.routeMap["UPLOAD"] = relativePath
}

// Modify delete image path.
func (s *ImageService) Delete(relativePath string) {
	s.routeMap["DELETE"] = relativePath
}

// Modify health check path.
func (s *ImageService) Healthcheck(relativePath string) {
	s.routeMap["HEALTHCHECK"] = relativePath
}

func (s *ImageService) registerAll() {
	if path, ok := s.routeMap["GET"]; ok {
		s.internal.GET(path, s.internal.factory.CreateGetImageHandler())
	}
	if path, ok := s.routeMap["UPLOAD"]; ok {
		s.internal.POST(path, s.internal.factory.CreateUploadImageHandler())
	}
	if path, ok := s.routeMap["Delete"]; ok {
		s.internal.DELETE(path, s.internal.factory.CreateDeleteImageHandler())
	}
	if path, ok := s.routeMap["HEALTHCHECK"]; ok {
		s.internal.GET(path, s.internal.factory.CreateHealthcheckHandler())
	}
}

func (s *ImageService) init() {
	s.routeMap = map[string]string{
		"GET":         "/image/*imagepath",
		"UPLOAD":      "/image",
		"DELETE":      "/image/*",
		"HEALTHCHECK": "/healthcheck",
	}
}
