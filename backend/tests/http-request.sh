curl http://localhost:3000$2 \
	--request $1 \
	--header "Content-Type: application/json" \
	--data "$3" \
| python3 -m json.tool
echo ''
