# REST API


## I. Get job list

1. HTTP

		GET /v1/jobs HTTP/1.1

2. Output Example

		{
			error: false  
			jobs : {
				1: {
					key: 'The key of the job (generated when create new job)',
					name: 'The name of job',
					status: 1,
					progress: 90,
					estimate: 25
				},

				2: {
                    key: 'The key of the job (generated when create new job)',
                    name: 'The name of job',
                    status: -1,
                    progress: 0,
                    estimate: 0
                }
                ...
            }
			}
		}
		


## II. Get job

1. HTTP

		GET /v1/job/<jobid> HTTP/1.1
		
2. Output Example

		{
			{
                key: 'The key of the job (generated when create new job)',
                name: 'The name of job',
                job_status: -1,
                progress: 0,
                estimate: 30
            }
		}


## III. Request job

1. HTTP

		PUT /v1/job HTTP/1.1
		
2. Input Example

		{
			name: 'Job name',
			action: 'create'
			work_load: {
				//TODO
			}
		}
		
3. Output Example

		{
			key: 'auto generated key',
			status: 1,
			message: ''
		}
		


## IV. Pause Job 

1. HTTP

		POST /v1/job/<key> HTTP/1.1
		
2. Input Example

		{
			action: 'pause'
		}
		
3. Output Example

		{
			status: 1,
			message: ''
		}
		
		
## V. Resume Job 

1. HTTP

		POST /v1/job/<key> HTTP/1.1
		
2. Input Example

		{
			action: 'resume'
		}
		
3. Output Example

		{
			status: 0,
			message: 'Example error message'
		}
		
		
## VI. Terminate (Cancel) Job 

1. HTTP

		DELETE /v1/job/<key> HTTP/1.1
		
2. Input Example

		{
			action: 'terminate'
		}
		
3. Output Example

		{
			status: 1,
			message: 'The job was terminated'
		}
		
		
		
## VI. Restart Job 

1. HTTP

		POST /v1/job/<key> HTTP/1.1
		
2. Input Example

		{
			action: 'restart'
		}
		
3. Output Example

		{
			status: 1,
			message: 'The job was restarted'
		}
		