import { useState, useEffect } from 'react';
import './App.css';
import {
  Container,
  Typography,
  Box,
  Grid,
  Card,
  CardContent,
  Link,
  Pagination,
} from '@mui/material';

function App() {
  const [data, setData] = useState(null);
  const [page, setPage] = useState(1);
  const jobsPerPage = 10;

  const host = import.meta.env.VITE_IS_DOCKER === 'true' ? 'http://localhost' : 'http://localhost:8080';

  const getJobs = async (page: number) => {
    try {
      const response = await fetch(`${host}/api/proxy-get`, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({
          url: `https://www.amazon.jobs/en/search.json?category%5B%5D=software-development&normalized_country_code%5B%5D=USA&radius=24km&industry_experience[]=less_than_1_year&facets%5B%5D=normalized_country_code&facets%5B%5D=normalized_state_name&facets%5B%5D=normalized_city_name&facets%5B%5D=location&facets%5B%5D=business_category&facets%5B%5D=category&facets%5B%5D=schedule_type_id&facets%5B%5D=employee_class&facets%5B%5D=normalized_location&facets%5B%5D=job_function_id&facets%5B%5D=is_manager&facets%5B%5D=is_intern&offset=${
            (page - 1) * 10
          }&result_limit=10&sort=relevant&latitude=&longitude=&loc_group_id=&loc_query=&base_query=&city=&country=&region=&county=&query_options=&`,
        }),
        cache: 'no-cache',
      });

      if (!response.ok) {
        throw new Error('Network response was not ok');
      }

      const rawData = await response.json();
      const responseData = JSON.parse(rawData.body);
      setData(responseData);
    } catch (error) {
      console.error('Error:', error);
    }
  };

  useEffect(() => {
    getJobs(page);
  }, [page]);

  const handleChangePage = (event, value) => {
    setPage(value);
  };

  const paginationStyles = {
    '& .MuiPaginationItem-previousNext': {
      color: 'white', // Change this to the desired color
    },
    '& .MuiPaginationItem-page': {
      color: 'white', // Change this to the desired color for the page numbers
    },
  };

  return (
    <Container
      maxWidth='sm'
      style={{ height: '100vh', display: 'flex', justifyContent: 'center' }}
    >
      <Box textAlign='center'>
        <Typography variant='h3' component='h1' gutterBottom>
          Amazon SDE 1 Jobs
        </Typography>
        <Typography variant='body1'>Total jobs: {data?.hits}</Typography>
        <Grid container spacing={2}>
          {data?.jobs?.map((job) => (
            <Grid item key={job.id} xs={12}>
              <Card
                style={{
                  height: '150px',
                  display: 'flex',
                  flexDirection: 'column',
                  justifyContent: 'center',
                }}
              >
                <CardContent>
                  <Typography variant='h6' component='div'>
                    {job.title}
                  </Typography>
                  <Typography variant='body2' color='textSecondary'>
                    {job.locations
                      .map((location) => {
                        const obj = JSON.parse(location);
                        return obj.location;
                      })
                      .join(', ')}
                  </Typography>
                  <Typography variant='body2' color='textSecondary'>
                    Updated time: {job.updated_time}
                  </Typography>
                  <Typography variant='body2' color='textSecondary'>
                    Posted date: {job.posted_date}
                  </Typography>
                  <Typography variant='body2'>
                    <Link
                      href={`https://www.amazon.jobs/en/jobs/${job.id_icims}`}
                      target='_blank'
                      rel='noopener'
                    >
                      Link to Post
                    </Link>
                  </Typography>
                </CardContent>
              </Card>
            </Grid>
          ))}
        </Grid>
        <Box display='flex' justifyContent='center' marginTop={2}>
          <Pagination
            count={Math.ceil(data?.hits / jobsPerPage)}
            page={page}
            onChange={handleChangePage}
            color='primary'
            sx={paginationStyles}
          />
        </Box>
      </Box>
    </Container>
  );
}

export default App;
