package constants

type ServiceConstantsV1 struct {
	Version  string   `json:"version"`
	Services Services `json:"services"`
}

type AmsOptions struct {
	TermsAndConditionsEventCode string `json:"termsAndConditionsEventCode"`
	TermsAndConditionsSiteCode  string `json:"termsAndConditionsSiteCode"`
	InstanceQuotaID             string `json:"instanceQuotaId"`
	TrialQuotaID                string `json:"trialQuotaId"`
}

type KafkaLimits struct {
	DefaultReplicas                 string `json:"default_replicas"`
	DefaultMinCleanbleRatio         string `json:"default_min_cleanble_ratio"`
	DefaultMinInsyncReplicas        string `json:"default_min_insync_replicas"`
	MinPartitions                   string `json:"min_partitions"`
	MaxPartitions                   string `json:"max_partitions"`
	DefaultSegmentTime              string `json:"default_segment_time"`
	DefaultMaxMessageTimestampDiff  string `json:"default_max_message_timestamp_diff"`
	DefaultSegmentIndexSize         string `json:"default_segment_index_size"`
	DefaultIndexIntervalSize        string `json:"default_index_interval_size"`
	DefaultLogSegmentSize           string `json:"default_log_segment_size"`
	DefaultDeleteRetentionTime      string `json:"default_delete_retention_time"`
	DefaultSegmentJitterTime        string `json:"default_segment_jitter_time"`
	DefaultFileDeleteDelay          string `json:"default_file_delete_delay"`
	DefaultMaximumMessageBytes      string `json:"default_maximum_message_bytes"`
	DefaultMessageTimestampType     string `json:"default_message_timestamp_type"`
	DefaultMinimumCompactionLagTime string `json:"default_minimum_compaction_lag_time"`
	DefaultFlushIntervalMessages    string `json:"default_flush_interval_messages"`
	DefaultFlushIntervalTime        string `json:"default_flush_interval_time"`
}
type Kafka struct {
	Ams    AmsOptions  `json:"ams"`
	Limits KafkaLimits `json:"limits"`
}

type ServiceRegistry struct {
	Ams AmsOptions `json:"ams"`
}

type Services struct {
	Kafka           Kafka           `json:"kafka"`
	ServiceRegistry ServiceRegistry `json:"serviceRegistry"`
}
