# Table: cloudflare_certificate_packs

This table shows data for Cloudflare Certificate Packs.

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id|String|
|zone_id|String|
|id (PK)|String|
|type|String|
|hosts|StringArray|
|certificates|JSON|
|primary_certificate|String|
|validation_records|JSON|
|validation_errors|JSON|
|validation_method|String|
|validity_days|Int|
|certificate_authority|String|
|cloudflare_branding|Bool|