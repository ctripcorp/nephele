# The Nephele Log Compatibility Standard

## Log Suite & Log Interface

Designed to be a compelete solution, Nephele provide a log suite compatible with [ES](https://www.elastic.co/products/elasticsearch) and [CAT](https://github.com/dianping/CAT). The log suite has two parts: develop kit and sync agent.

The log develop kit is a go package providing nephele log programming interface, abbreviated as NLPI. Calling NLPI from your code forces your program to dump log files within some given constraints: specific alignment, particular layout, mangling names, etc. All these details are called the Nephele Log Binary interface, or NLBI. And when the moment is ripe, a background program will collect and decompose theses dumped log files. And that is the sync agent. It translates the content of log files into corresponding log units(CAT message or ES document or something similar) and send them to the target servers.

**Keep reading through, you will learn about NLBI. Yet NLPI is documented in The Nephele Programming Interface() in detail, not here.**

## Compatibility

The NLBI compatibility is defined with the equation above: given a log file generated with NLBI, it is able to be decomposed into single logs and every single log contains adequate infomation to be recognized as a given log unit.

Currently NLBI is compatible with ES and CAT.

## Versioning

The NLBI and the Nephele program are independant from versioning. 

The NLBI version consists of two numbers connected with a dot. 

This document(The Nephele Log Compatibility Standard) and The NLBI share a concert version number.

## Allowed Changes

The following will cause the minor version number to increase, say from "1.0" to "1.1".

## Prohibited Changes

The following will cause the major version number to increase, say from "0.9" to "1.0".

## Implementation

## FAQ

## Contact
