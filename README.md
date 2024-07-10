## Amazon SDE 1 Jobs

simple web app that show all US amazon software engineer job that have less than 1 year experience (SDE 1)

I built this because sometimes the yeah of experince filter on https://www.amazon.jobs/en/search is missing (I think it is something that do with the cookies)

I interviewed with Amazon for SDE2 but I failed and got down level to SDE1.

Amazon don't really have a lot of SDE1 open in my area and my interview result only valid for 12 months so I need to check this everyday and hopefully snipe some job xd.

The reality is that very likely I will not find a SDE1 position in my area in the next 12 months but I like copium.

### How to run

You can run with docker or without docker which mean you either need to have docker installed or have go, node, vite, npm ready installed.

You also need to create a .env file in the backend folder and put `ANALYTICS_ID=XXXX`

ANALYTICS_ID is one of the auto generate cookies when visit the amazon job site

I used postman to make a get request on https://www.amazon.jobs/en/search.json?category%5B%5D=software-development&normalized_country_code%5B%5D=USA&radius=24km&industry_experience[]=less_than_1_year&facets%5B%5D=normalized_country_code&facets%5B%5D=normalized_state_name&facets%5B%5D=normalized_city_name&facets%5B%5D=location&facets%5B%5D=business_category&facets%5B%5D=category&facets%5B%5D=schedule_type_id&facets%5B%5D=employee_class&facets%5B%5D=normalized_location&facets%5B%5D=job_function_id&facets%5B%5D=is_manager&facets%5B%5D=is_intern&offset=20&result_limit=10&sort=relevant&latitude=&longitude=&loc_group_id=&loc_query=&base_query=&city=&country=&region=&county=&query_options=& then it auto generate cookies for me

Run with docker:
- `docker-compose up` then it should work on `localhost:80`

Run without docker:
- cd into backend folder and run `go run server.go`
- cd into frontend folder and run `npm run dev`
- then it should work on `localhost:5173`


### What it Look like
![Screenshot 2024-07-09 233536](https://github.com/LonelyLok/amazon-sde-1-job-search/assets/40349145/efcb3d3d-99ce-4723-b067-259c44e15d91)
