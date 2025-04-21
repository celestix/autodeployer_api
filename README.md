# AutoDeployer Backend

## API Structure

```
/api
    /auth
        /login      POST => LOGIN USER               [  LOGIN VIA  ]
        /signup     POST => ONE TIME ROUTE (SIGNUP)  [   GITHUB    ]
    /project
        /list       GET  => LIST ALL PROJECTS
        /new        POST => CREATE NEW PROJECT
        /info       GET  => A PROJECT INFO
        /graph      GET  => PER PROJECT USAGE
        /deploy     POST
        /resources  GET  => CPU, RAM, NETWORK usage
        /advanced   GET  => ADVANCED INFO (NO. OF DEPLOYS, LOGS)
    /dashboard
        /graph      GET  => OVERALL RESOURCE USAGE
    /hook           POST => WEBHOOK ENDPOINT FOR GITHUB

```