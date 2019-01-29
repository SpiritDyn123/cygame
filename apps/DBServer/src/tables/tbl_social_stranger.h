#pragma once

#include "../core/dbtable.h"

class TblSocialStranger
	: public DBTable<100, 100>
{
public:
	using DBTable::DBTable;
	//TblSocialStranger(SQLConnection& conn, const char* db, const char* tbl);
	// ��¼İ���˷�������ʱ��
	int AddStranger(uint32_t stranger_pid);
	// ��ȡָ��ʱ����ڵ�İ���˷�������
	int GetStrangerCnt();
	// ��¼�Լ����µĲ�ѯʱ��
	int UpdateCheckTime();
	// ��ȡ�Լ����µĲ�ѯʱ��
	int GetCheckTime();
};

