# Book_Api

# This simple Api app in golang

1. POST request try this command

    curl localhost:8080/books --include --header "Content-Type:application/json" -d @body.json --request "POST"
2. GET request try this command

    1. curl localhost:8080/books
    2. curl localhost:8080/books/id
3. PATCH request try this command

    1. curl localhost:8080/checkout?id=2 --request "PATCH"
    2. curl localhost:8080/return?id=2 --request "PATCH
