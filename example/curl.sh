
## When using Curl, you call API like this 

# Register
curl http://localhost:8000/api/auth/register \
	-H "Content-Type: application/json" \
	--request POST \
	-d '{"email":"cong2@gmail.com","password":"123456Aa","confirm_password":"123456Aa"}' 

# Login
curl http://localhost:8000/api/auth/login \
	-H "Content-Type: application/json" \
	--request POST \
	-d '{"email":"cong@gmail.com","password":"123456Aa"}' 

# Get info - replace TOKEN with your token
curl http://localhost:8000/api/auth/info \
	-H "Content-Type: application/json" \
	-H "Authorization: ${TOKEN}" \
	--request GET 
