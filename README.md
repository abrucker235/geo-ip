# Ideas
* Implement redis cache with secondary kubernetes process on cron job to keep cache up-to-date once a week.

# Extensions
* Expand the capabilities to look up geo location per IP Address. This capability could be used for gathering metrics around where API interactions are coming from which can be used for various analytics.
    * Determining potentially future Data Centers or Availability Zones to use.
    * Build graph database showing relationship between what IPs call and which API Resources they are trying to call.