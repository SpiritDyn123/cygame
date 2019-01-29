#pragma once

#include "../core/dbtable.h"

class TblLotteryRecord
	: public DBTable<100, 100>
{
public:
	using DBTable::DBTable;
	// ������ҳ鿨��¼
	int AddLotteryRecord(uint32_t type, std::string& data);
	// ��ȡ��ҳ鿨��¼
	int GetLotteryRecords(uint32_t type, uint32_t index, uint32_t limit, uint32_t time_beg, uint32_t time_end, 
		const std::function<void(std::string& data)>& fn);
};

