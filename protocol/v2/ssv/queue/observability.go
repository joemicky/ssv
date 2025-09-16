package queue

import (
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/metric"

	"github.com/ssvlabs/ssv/observability"
	"github.com/ssvlabs/ssv/observability/metrics"
)

const (
	observabilityName      = "github.com/ssvlabs/ssv/protocol/v2/ssv/queue"
	observabilityNamespace = "ssv.queue"
)

var (
	meter = otel.Meter(observabilityName)

	// ValidatorQueuesInboxSizeMaxMetric represents the max queue size across all validator-related message queues
	// since SSV node start.
	ValidatorQueuesInboxSizeMaxMetric = metrics.New(
		meter.Int64Gauge(
			observability.InstrumentName(observabilityNamespace, "validator_queues_inbox_size_max"),
			metric.WithUnit("{size}"),
			metric.WithDescription("max size validator queues reached since node start")),
	)

	// CommitteeQueuesInboxSizeMaxMetric represents the max queue size across all committee-related message queues
	// since SSV node start.
	CommitteeQueuesInboxSizeMaxMetric = metrics.New(
		meter.Int64Gauge(
			observability.InstrumentName(observabilityNamespace, "committee_queues_inbox_size_max"),
			metric.WithUnit("{size}"),
			metric.WithDescription("max size committee queues reached since node start")),
	)
)
