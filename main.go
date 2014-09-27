package main

import (
	"database/sql"
	"fmt"
	_ "github.com/denisenkom/go-mssqldb"
)

func main() {
	db, err := sql.Open("mssql", "server=192.168.0.4;user id=UCifecGest;database=CifecGEST;password=Infor?*eurO;")
	if err != nil {
		fmt.Printf("Couldn't open DB: %s", err)
		return
	}
	fmt.Println("Conected")
	st, err := db.Prepare(`SELECT     ventas_cabecera.cod_cliente, ventas_cabecera.nombre,ventas_linea.albaran, ventas_linea.num_caja, ventas_linea.cod_articulo, ventas_linea.descripcion
FROM         ventas_linea INNER JOIN
                      ventas_cabecera ON ventas_linea.cod_empresa = ventas_cabecera.cod_empresa AND ventas_linea.albaran = ventas_cabecera.albaran AND 
                      ventas_linea.tipo = ventas_cabecera.tipo
WHERE     (ventas_linea.tipo = 'A') AND (ventas_linea.importe_iva = 0) AND (ventas_linea.cod_articulo <> 'NOTA') OR
                      (ventas_linea.tipo = 'A') AND (ventas_linea.cod_articulo = 'PENDENT')`)

	if err != nil {
		fmt.Printf("Couldn't Prepare ST: %s", err)
		return
	}
	rows, err := st.Query()
	if err != nil {
		fmt.Printf("Couldn't query DB: %s", err)
		return
	}

	for rows.Next() {
		cli := ""
		nam := ""
		alb := 0
		caj := 0
		cod := ""
		desc := ""
		rows.Scan(&cli, &nam, &alb, &caj, &cod, &desc)
		fmt.Println(format(cli, 5), format(nam, 40), alb, caj, format(cod, 10), desc)
	}
	rows.Close()

	defer db.Close()
}

func format(str string, n int) string {

	for i := len(str); i < n; i++ {
		str += " "
	}
	return str
}
