
## When using Gotman, you call API like this 

# Register
gotman -u http://localhost:8000/api/auth/register \
-m POST \
-h 'Content-Type: application/json' \
-b '{"email":"cong3@gmail.com","password":"123456Aa","confirm_password":"123456Aa"}' 

# Login
gotman -u http://localhost:8000/api/auth/login \
-m POST \
-h "Content-Type: application/json" \
-b '{"email":"cong@gmail.com","password":"123456Aa"}' 

# Get info - replace TOKEN with your token
gotman -u http://localhost:8000/api/auth/info \
-m GET 
-h "Content-Type: application/json" \
-h "Authorization: ${TOKEN}" \

## Load from config file 
gotman -u '/auth/login' \
-f example/config.json \
-m POST \
-b '{"email":"cong@gmail.com","password":"123456Aa"}' \ 
