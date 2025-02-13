# 🏢 go-practice
### 🛠️ Tools
Go Language

### 📚 Go Libraries
* Google Go UUID
* **pq**: Pure Go postgres driver for Go's database/sql package

## 📋 Briefing
go language capabilities were explored to evaluate the integration requirements of building a functioning web app.
### Workflow Considerations
1. Connection should be established to local PostgreSQL and successful output message should generate when connection is established.
   - If connection is attempted but isn't made to DB, then an error message should be generated.
2. Additional entries to the DB's table should produce a new UUID.
3. Duplicate queries should be deleted from the DB's table. Success output message will be generated upon deletion of duplicate entries. 
   -  Error output message will be generated if there's an unsuccessful deletion of any duplicate entries.
4. DB connection will close after performing the necessary actions.
