package gmagick

/*
#include <wand/wand_api.h>
*/
import "C"

type MetricType int

const (
	METRIC_UNDEFINED                          MetricType = C.UndefinedMetric
	METRIC_MEAN_ABSOLUTE_ERROR                MetricType = C.MeanAbsoluteErrorMetric
	METRIC_MEAN_SQUARED_ERROR                 MetricType = C.MeanSquaredErrorMetric
	METRIC_PEAK_ABSOLUTE_ERROR                MetricType = C.PeakAbsoluteErrorMetric
	METRIC_PEAK_SIGNAL_TO_NOISE_RATIO         MetricType = C.PeakSignalToNoiseRatioMetric
	METRIC_ROOT_MEAN_SQUARED_ERROR            MetricType = C.RootMeanSquaredErrorMetric
)
