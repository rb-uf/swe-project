sh http-request.sh POST create-subject '{ "Name": "Marston" }'
sh http-request.sh POST create-review '{ "Subject": "Marston", "Rating": 5, "Text": "It is alright.", "Author": "name", "AuthorID": 5678 }'
sh http-request.sh GET get-subject-reviews '{ "Name": "Marston", "MaxReviews": 10000 }'
