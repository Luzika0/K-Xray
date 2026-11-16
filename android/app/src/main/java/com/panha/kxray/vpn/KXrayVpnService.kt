package com.panha.kxray.vpn

import android.net.VpnService
import android.content.Intent
import android.util.Log

/**
 * Platform lifecycle boundary for the Android VPN.
 *
 * The engine is deliberately not started here yet. A future implementation must:
 * 1. Obtain explicit user consent with VpnService.prepare().
 * 2. Establish the TUN interface and configure routes/DNS.
 * 3. Pass the descriptor to the selected engine adapter.
 * 4. Protect the engine's transport socket to prevent routing loops.
 * 5. Tear down the engine and descriptor on revoke/stop.
 */
class KXrayVpnService : VpnService() {
    override fun onStartCommand(intent: Intent?, flags: Int, startId: Int): Int {
        Log.i(TAG, "VPN service scaffold started; no tunnel engine is integrated.")
        stopSelf(startId)
        return START_NOT_STICKY
    }

    override fun onRevoke() {
        Log.i(TAG, "VPN permission revoked.")
        super.onRevoke()
        stopSelf()
    }

    companion object {
        private const val TAG = "KXrayVpnService"
    }
}
