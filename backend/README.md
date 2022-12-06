Mortgage Calculator
===========

This api will calculate the scheduled payment for any given mortgage input, those input have this schema:

```
{
  "anual_interest": 0, 
  "downpayment": 0, # More than 5%
  "period": 0, # 5 year increments between 5 and 30 years
  "price": 0,
  "schedule": "string" # accelerated bi-weekly | bi-weekly | monthly
}
```

You can clone the repository and run the project using:

```
git clone https://github.com/highercomve/mortgage_calculator.git
cd mortgage_calculator
go run .
```

After that you can test the api by calling the api like this:

```
curl --request POST \
  --url http://localhost:9090/api/v1/calculate \
  --header 'Content-Type: application/json' \
  --data '{
	"schedule": "monthly",
	"price": 300000,
	"downpayment": 40000,
	"period": 20,
	"anual_interest": 5
}'
```

In case of success you will have an response like:

```
{
	"schedule": "monthly",
	"number_of_payments": 240,
	"payment": 1769.077354544168
}
```

If the service fail you will get an:

```
{
	"message": "period should be 5 year increments between 5 and 30 years"
}
```

For more documentation you can go to:

http://localhost:9090/swagger/index.html

