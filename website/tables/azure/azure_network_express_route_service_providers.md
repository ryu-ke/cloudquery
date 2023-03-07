# Table: azure_network_express_route_service_providers

https://learn.microsoft.com/en-us/rest/api/expressroute/express-route-service-providers/list?tabs=HTTP#expressrouteserviceprovider

The composite primary key for this table is (**subscription_id**, **id**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|subscription_id (PK)|String|
|id (PK)|String|
|location|String|
|properties|JSON|
|tags|JSON|
|name|String|
|type|String|