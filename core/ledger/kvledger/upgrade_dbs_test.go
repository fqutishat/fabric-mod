/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package kvledger

// TODO: Re-enable test after upgrade is implemented in fabric-peer-ext
//func TestUpgradeWrongFormat(t *testing.T) {
//	conf, cleanup := testConfig(t)
//	conf.HistoryDBConfig.Enabled = false
//	defer cleanup()
//	provider := testutilNewProvider(conf, t, &mock.DeployedChaincodeInfoProvider{})
//
//	// change format to a wrong value to test UpgradeFormat error path
//	err := provider.idStore.Put(formatKey, []byte("x.0"))
//	provider.Close()
//	require.NoError(t, err)
//
//	err = UpgradeDBs(conf)
//	expectedErr := &dataformat.ErrFormatMismatch{
//		ExpectedFormat: dataformat.PreviousFormat,
//		Format:         "x.0",
//		DBInfo:         fmt.Sprintf("leveldb for channel-IDs at [%s]", LedgerProviderPath(conf.RootFSPath)),
//	}
//	require.EqualError(t, err, expectedErr.Error())
//}
