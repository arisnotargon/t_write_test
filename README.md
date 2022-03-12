
# Topic 2 monitor
## 1
Prepare a health check api (or name it "ping"),then we can check whether the service is in normal state by requesting this api.
## 2
Write logs  when an exception occurs.

## 3
Use Prometheus and Grafana  to monitor the state of our application.


<br>
<br>
# Topic 3 optimize:
## 1
use databases for durability backup car sold data.when the program has been terminated unexpectedly,we can restore our data after restarted it. 
## 2
If we update the  rate every minute(Hint by "R is calculated based on sliding time window, which moves once a minute"),we well lost the data befor update where the api "buffer" has been requested.So we Can use relational database to solve this problem.For example, use the following sql: 

```sql
select count(*) from `car_sold_log` where car_id > 1647075998602136221
```
to get the num of sold cars number in recent one hour,1647075998602136221 is the result of the time which is an hour before the current time convert to UNIX nanosecond timestamp.
