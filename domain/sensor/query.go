package sensor

const CreateSensorDataQuery = `
	INSERT INTO sensor_data (id, device_id, temperature, humidity)
	VALUES (:id, :device_id, :temperature, :humidity)
`

const CreateDeviceControlQuery = `
	INSERT INTO device_commands (id,device_id, command, status)
	VALUES (:id, :device_id, :command, :status)
`

const GetSensorDataByDeviceIDQuery = `
	SELECT device_id, temperature, humidity, timestamp
	FROM sensor_data
	WHERE device_id = :device_id
	ORDER BY timestamp DESC
	LIMIT 10
`

const GetDeviceControlByDeviceIDQuery = `
	SELECT id, device_id, command, status, created_at
	FROM device_commands
	WHERE device_id = :device_id
	ORDER BY created_at DESC
	LIMIT 10
`
