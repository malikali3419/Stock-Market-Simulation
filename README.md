# Stock Market Simulations Software

This Git repository is a comprehensive stock market simulation application. It provides a set of RESTful API endpoints for managing users, stock data, and transactions, with the goal of simulating stock market activities.

## Table of Contents

- [Installation](#installation)
- [User Authentication](#user-authentication)
- [Working](#Working)
- [License](#license)

## Installation


1. Clone the repository: git clone https://github.com/malikali3419/Stock-Market-Simulation.git
2. Change to the project directory:
```bash
$ cd Stock Market Simulation/
````
3. Install dependencies:
```bash
$ go get -d -v $(cat dependencies.txt)
```
5. Set up the environment variables:
- `DATABASE_URL`: Connection string to your PostgreSQL database
- `PORT`: Add your desire PORT
6. Migrate the models ``` go run migrate/migrate.go```
7. Get the build of project:
```bash
$ go build -o Stock Market Simulation
```
8. Run the project:
```bash
$ go run main.go
```
## User Authentication

1. Sign up: Send a POST request to `/signup` with the following JSON body:
   `{
   "username": "your-username",
   "password": "your-password"
   }`
2. Log in: Send a POST request to `/login` with the same JSON body. You'll receive a JWT token in the response.
## Working

1. Authenticate: Include the JWT token in the `Authorization` header for subsequent requests.

2. Add Stocks: Send a POST request to `/stocks` with the following JSON body:
   `{
   "ticker":"GOOGL",
   "openprice":10,
   "closeprice":20,
   "high":50,
   "low":20,
   "volume":900
   }`
   Get all stocks: Send a GET request to `/stocks` to retrieve a list of all Stocks.

4. To get all users send Get request on `/allUsers` .
5. To get a Specific user send get request or this url `/user/name`. 
6. To get the Stock of specific ticker send Get request on this url `/stock/ticker name`
7. Do Transaction by sending POST request on this url `/transactions` this will be run on background
8. To get all transactions send Get request on this url `/transactions/:user_id` this will return all the transaction of the specific user
9. To get the transaction of the user between specific time send get request on this url `/transactions/:user_id/:start_time/:end_time`

## Contributing

1. Fork the repository.

2. Create a new branch: `git checkout -b feature-new-feature`
3. Make your changes and commit: `git commit -m "Add new feature`
4. Push to the branch: `git push origin feature-new-feature`
5. Submit a pull request.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
