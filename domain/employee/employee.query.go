package employee

const (
	queryStore = `insert into employees (id, first_name, last_name, title, work_phone) values (:id, :first_name, :last_name, :title, :work_phone)`
)

const (
	queryGetIDByWorkPhone = `select id from employees where work_phone = ? and is_deleted = 0`
)
