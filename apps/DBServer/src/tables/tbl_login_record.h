#pragma once

#include "../core/dbtable.h"

class TblLoginRecord
	: public DBTable<100, 100>
{
public:
	using DBTable::DBTable;
	// ����������߼�¼
	int AddLoginRecord(std::string ip);
	// ��ȡ������߼�¼
	int GetLoginRecords(uint32_t index, uint32_t limit, const std::function<void(std::string ip, uint32_t time_stamp)>& fn);
};

