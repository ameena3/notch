<html>
    <head>
    <title>{{.Title}}</title>
    </head>
    <h2>Enter the product that you want to remove from the cart</h2>
    <body>
        <form action="/home?cart_id={{.CartID}}" method="post">
            <table>
                <tr>
                <th>Name</th>
                <th>Description</th>
                <th>Price</th>
                <th>Remove</th>
                </tr>
                {{range .Products}}
                <tr>
                    <td>{{.Name}}</td>
                    <td>{{.Description}}</td>
                    <td>{{.Price}}</td>
                    <td><input type="checkbox" name={{.SKU}} value={{.SKU}}></td>
                {{end}}
            </table>
            <input type="submit" value="Submit">
    </form>
    </body>
</html>