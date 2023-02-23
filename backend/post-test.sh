curl http://localhost:3000/create-subject \
	--request POST \
	--header "Content-Type: application/json" \
	--data-urlencode @json-samples/subject-example.json
