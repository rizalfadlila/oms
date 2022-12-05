package initiator

import (
	"github.com/jatis/oms/init/service"
	"github.com/jatis/oms/lib/database/sql"
)

func (i *initiator) newBasic() {
	i.basic = &service.Basic{
		MariaClient: i.NewMariaClient(),
	}
}

func (i *initiator) NewMariaClient() *sql.Store {
	return sql.New(i.config.Database.MariaDB)
}
