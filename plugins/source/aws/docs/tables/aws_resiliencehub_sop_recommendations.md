# Table: aws_resiliencehub_sop_recommendations

https://docs.aws.amazon.com/resilience-hub/latest/APIReference/API_SopRecommendation.html

The composite primary key for this table is (**account_id**, **region**, **recommendation_id**).

## Relations

This table depends on [aws_resiliencehub_app_assessments](aws_resiliencehub_app_assessments.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id (PK)|String|
|region (PK)|String|
|recommendation_id (PK)|String|
|reference_id|String|
|service_type|String|
|app_component_name|String|
|description|String|
|items|JSON|
|name|String|
|prerequisite|String|