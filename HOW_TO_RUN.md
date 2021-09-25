1. Clone the repo 
```bash
git clone https://github.com/CreamyMilk/buycoinz.git
```

2. Make a copy  of `.env.sample` then rename it to `.env` & update the environment variables
```bash
cp .env.sample .env
```

3. Run the app using

```bash
go run main.go
```


#### Environment Variables Descriptions
<pre>
* ENV  - The current deployment mode ("staging",production)
* PORT - port to run the server on
* PAYSTACK_PROD_KEY - A valid paystack production API key
* PAYSTACK_TEST_KEY - A valid paystack Test API key 
</pre>

#### Testing Environment variables
<pre>
* TEST_BANKCODE   - A legitimate Bank code 
* TEST_ACCOUNTNO  - A legitimate accountNo 
* TEST_VALID_ACCOUNT_NAME    - The actual account name that corresponds with the mentioned (AccoutNo,BankCode) pair
* TEST_PREFERED_ACCOUNT_NAME - User specified Custom name to be persited for testing purposes
</pre>


## Testing 
Each of the respective packages has corresponding test in their folders

```bash
go test -v ./services/      # Some test cases require valid test api tokens to run
go test -v ./paystack/      # (API intrgrations Tests you need a valid token to run this tests)
go test -v ./utils/         
```


### Assumptions Made
* Account Numbers can vary in length and they can contain leading zeros hence 
accounts like (001230) and( 1230 ) are different.
* All account number are unique irrespective of the bank of origin
* Aspects like authorization ,rate-limiting and CSRF are beyond the scope of the task
* Authentication and user management was beyound the scope iof the task.

<br />

#### Why I don't use an actual DB for my implmentation
> I resorted to using an in memory hash map to store the prefered  
> account name.
> The inclusion of a database would have been beyond the 
> scope of the task due to the creation of a user model and 
> creation of a mock database to run my tests.