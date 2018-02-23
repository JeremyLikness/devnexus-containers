# Docker Compose

`cat docker-compose.yaml`

In the root directory:

`docker-compose up`

Navigate to [http://localhost:8000](http://localhost:8000)

Then, in another window:

`docker-compose scale goservice=2`

Increase from 2 to 3, etc. and watch the changes.

Then set back to 1.