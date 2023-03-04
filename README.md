URL shortener projects that implements hexagonal architecture pattern. 

App was written in Golang with chi web framework. It supports both mongoDb and redis database.

Usage: 
    - To shorten an url use cmd:
        curl -XPOST localhost:8000/ -H 'Content-Type: application/json' -d '{"url":"https://www.youtube.com/watch?v=Shh2FYRQAlY"}'
    which return in response:
        
        {"code":"eQEwsUb4g","url":"https://www.youtube.com/watch?v=Shh2FYRQAlY","created_at":1677922449}
    
    we can then use the code to redirect to the page, just paste in the browser:
        
        localhost:8000/eQEwsUb4g or use curl -XGET  curl -XPOST localhost:8000/eQEwsUb4g
