<html>
    <head>
    <title>{{.Title}}</title>
    </head>
    <h2>Enter the product that you want to purchase</h2>
    <body>
        <form action="/cart?cart_id={{.CartID}}" method="post">
            <table>
                <tr>
                <th>Name</th>
                <th>Description</th>
                <th>Price</th>
                <th>Puchase</th>
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