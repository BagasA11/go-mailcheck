
# Email checker

An email checker with Go language

## Features

- check email format using html form
- check email format using json request

## Appendix

This application is created using the package from 

 🔗 Links
https://github.com/badoux/checkmail



## Run Locally

Install go compiler. 

 🔗 Links

https://go.dev/dl/



Clone the project

```bash
  git clone https://github.com/BagasA11/go-mailcheck.git
```

then, you will get folder named like this:

`\go-mailcheck`

Move project to Go compiler folder

`D:\go\`

Open terminal or other CLI, (ig: cmd)
then, go to the project directory

```bash
cd  D:\go\go-mailcheck
```

Start the server

```bash
  go run main.go
```


## Usage/Examples

```bash
<!DOCTYPE html>
<html>
<head>
    
</head>
<body>
    
    <form id="form" action="/api/email/json" method="post">
        <label for="email">email:</label>
        <input type="email" id="email" name="email" placeholder = "" required ><br><br>
        <input type="submit" value="submit" >
    </form>

    <script>
        document.getElementById("form").addEventListener("submit", function(event) {
            event.preventDefault(); // prevent sending form

            // bind email from html form
            const email = document.getElementById("email").value;
            

            // create form object
            const formData = {
                email: email
            };

            // send post request to endpoint
            fetch("/api/email/json", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json"
                },
                body: JSON.stringify(formData)
            })
            .then(response => response.json())
            .then(data => {
                console.log("Response:", data);
                // response handling
            })
            .catch(error => {
                console.error("failed send POST request:", error);
            });
        });
    </script>
</body>
</html>


```


## Environment Variables

You can configure this application on the env file as needed

- create `.env` file on root folder
- copy all `.env.example` to `.env` file
- Change the attributes on the '.env' as per your needs

Here are the available attributes:
- endpoint
An endpoint is the URL to access the service. By default, the endpoint values may look like the following values:

`ENDPOINT=email`

- port
Port is the location where the server application is running. By default, the application runs on port 8080. You can override the following attributes:

`PORT={your_port}`


## API Reference

### validate email format with raw html

```http
  POST localhost:{PORT}/api/{ENDPOINT}
```

#### header

| key | value     | |
| :-------- | :------- | :------------------------- |
| `Content-Type` | `multipart/form-data` ||


#### body

| key | value     | |
| :-------- | :------- | :------------------------- |
| `email` | `someone@gmail.com` ||


### validate email format with json

```http
  POST localhost:{PORT}/api/{ENDPOINT}/json
```

#### header

| key | value     | |
| :-------- | :------- | :------------------------- |
| `Content-Type` | `application/json` ||


#### body

```bash
{
    "email":"someone@gmail.com"
}
```