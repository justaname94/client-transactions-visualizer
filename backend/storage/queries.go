package storage

const (
	// BuyerInfo returns several informations related to a buyer on the database.
	// The information returned can be summarized in overall user information,
	// other buyers that used the same ip and a list of recommended products.
	// The recommended products work by traversing the transactions of all the
	// other buyers that bought the same products of the user and sorting them
	// by amount.
	// Ex: If the buyer bought pizza, then it will look at all the other
	// transactions that contain pizza and will store what other things they
	// bought (ex: soda), after that sorts them by amount and recommends which
	// items where bought the most along with 'pizza' and all the other products
	// in the buyer transactions. (and now I want pizza)
	BuyerInfo = `query UserInfo($id: string) {
		buyer(func: eq(id, $id)) {
			name
			age
			id
			
			transaction {
				id
				device
				ip as ip
				product {
					name
					price as price
				}
				total : sum(val(price))
			}
		}
			
		relatedIpBuyers(func: eq(ip, val(ip)), first: 5) @filter(NOT uid(ip))  
		@normalize {
			ip : ip
			~transaction  {
				name: name
				id:  id
			}
		}
	
		
		 var(func: eq(id, $id))  {
				transaction {
			times as math(1)
					product  {
						~product  {
							product {
							ocurr as math(times)
								product_ids as id
							name 
							}
						}
					}
				}
		}
	
		var(func: eq(id, $id)) {
			transaction {
				product {
					user_products as id
				}
			}
		}
	
		recommendations(func: uid(product_ids), orderdesc: val(ocurr), first:10) 
		@filter(NOT uid(user_products)) {
			name
			price
		}
	}`

	// AllBuyers return all the buyers on the database
	AllBuyers = `
		query AllBuyers{
			buyers(func: type(Buyer)) {
				id
				name
				age
			}
		}
	`

	// AllBuyersPaginated return all the buyers on the database paginated
	AllBuyersPaginated = `
		query AllBuyersPaginated($page: int, $limit: int) {
			buyers(func: type(Buyer), offset: $page, first: $limit) {
				id
				name
				age
			}
		}
	`

	// AllProducts return all the products on the database.
	AllProducts = `
	query AllProducts{
		products(func: type(Product)) {
			id
			name
			price
		}
	}
	`

	// GetDate gets a date from the DB
	GetDate = `
	query GetDate($date: string) {
		date(func: eq(date, $date)) {
			date
		}
	}
	`
)
