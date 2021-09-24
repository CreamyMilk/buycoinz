1. Clone the repo 
```bash
git clone https://github.com/CreamyMilk/buycoinz.git
```

2. Copy .env.sample as .env & update the environment variable 
```bash
cp .env.sample .env
```

3. Run the app using 
```bash
go run main.go
```


#### Environment Varibles
<pre>
* ENV  - The current deployment mode ("staging",production)
* PORT - port to run the server
* PAYSTACK_PROD_KEY - A valid paystack production api key
* PAYSTACK_TEST_KEY - A valid paustack test api key 
</pre>

#### Testing Enironment varibles
<pre>
* TEST_BANKCODE   - A legitimate Bank code 
* TEST_ACCOUNTNO  - A legitimate accountNo 
* TEST_VALID_ACCOUNT_NAME    - The actual account name that corresponds with the mentioned (AccoutNo,BankCode) pair
* TEST_PREFERED_ACCOUNT_NAME - User specied Custom name to be persited for testing purposes
</pre>


## Testing 
Each of the respective packages has corresponding test in their folders

```bash
go test -v ./services/      # Some test cases require valid test api tokens to run
go test -v ./paystack/      # (API intrgrations Tests you need a valid token to run this tests)
go test -v ./utils/         
```


### Assumptions Made
* Account Number can vary in length and the can contain leading zeros hene accounts like (00123) and( 123 ) are diffent numbers
* All account number are unique in respective of the bank of origin
* Aspects like authroization ,rate-limiting and CSRF are beyond the scope of the task


### Why we don't use an actual DB 
>  I resorted to using an in memory hash map to store the prefered  
>  account name details as inclusion of a database would have been beyond the 
>  scope of the task