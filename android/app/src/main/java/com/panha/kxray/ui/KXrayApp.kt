package com.panha.kxray.ui

import androidx.compose.foundation.layout.Arrangement
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.Row
import androidx.compose.foundation.layout.Spacer
import androidx.compose.foundation.layout.fillMaxSize
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.height
import androidx.compose.foundation.layout.padding
import androidx.compose.foundation.shape.RoundedCornerShape
import androidx.compose.material3.Button
import androidx.compose.material3.Card
import androidx.compose.material3.CardDefaults
import androidx.compose.material3.MaterialTheme
import androidx.compose.material3.OutlinedButton
import androidx.compose.material3.Scaffold
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.runtime.getValue
import androidx.compose.runtime.mutableStateOf
import androidx.compose.runtime.remember
import androidx.compose.runtime.setValue
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.text.font.FontWeight
import androidx.compose.ui.unit.dp

private enum class EngineMode { XRAY, PING_TUNNEL }

@Composable
fun KXrayApp() {
    var selectedMode by remember { mutableStateOf(EngineMode.XRAY) }

    Scaffold { padding ->
        Column(
            modifier = Modifier
                .fillMaxSize()
                .padding(padding)
                .padding(horizontal = 24.dp, vertical = 28.dp),
            verticalArrangement = Arrangement.spacedBy(20.dp)
        ) {
            Column(verticalArrangement = Arrangement.spacedBy(4.dp)) {
                Text(
                    text = "k-Xray",
                    style = MaterialTheme.typography.headlineLarge,
                    fontWeight = FontWeight.Bold
                )
                Text(
                    text = "Native VPN client",
                    style = MaterialTheme.typography.bodyMedium,
                    color = MaterialTheme.colorScheme.onSurfaceVariant
                )
            }

            Card(
                modifier = Modifier.fillMaxWidth(),
                shape = RoundedCornerShape(24.dp),
                colors = CardDefaults.cardColors(
                    containerColor = MaterialTheme.colorScheme.surfaceContainer
                )
            ) {
                Column(
                    modifier = Modifier.padding(20.dp),
                    verticalArrangement = Arrangement.spacedBy(12.dp)
                ) {
                    Text("CONNECTION", style = MaterialTheme.typography.labelMedium)
                    Text(
                        "Not connected",
                        style = MaterialTheme.typography.headlineMedium,
                        fontWeight = FontWeight.SemiBold
                    )
                    Text(
                        "VPN engines are not integrated in this foundation build.",
                        color = MaterialTheme.colorScheme.onSurfaceVariant,
                        style = MaterialTheme.typography.bodyMedium
                    )
                    Button(
                        onClick = { /* Intentionally disabled until an engine is integrated. */ },
                        enabled = false,
                        modifier = Modifier.fillMaxWidth()
                    ) {
                        Text("Connect — coming in a later phase")
                    }
                }
            }

            Text("SELECT ENGINE", style = MaterialTheme.typography.labelLarge)
            EngineOption(
                title = "Xray Core",
                subtitle = "VLESS · VMess · Trojan · Shadowsocks",
                selected = selectedMode == EngineMode.XRAY,
                onClick = { selectedMode = EngineMode.XRAY }
            )
            EngineOption(
                title = "PingTunnel",
                subtitle = "ICMP transport · direct TUN mode planned",
                selected = selectedMode == EngineMode.PING_TUNNEL,
                onClick = { selectedMode = EngineMode.PING_TUNNEL }
            )

            Spacer(modifier = Modifier.height(4.dp))
            OutlinedButton(
                onClick = { /* Profile import UI is a later milestone. */ },
                modifier = Modifier.fillMaxWidth()
            ) {
                Text("Profiles & import (planned)")
            }

            Row(
                modifier = Modifier.fillMaxWidth(),
                horizontalArrangement = Arrangement.Center,
                verticalAlignment = Alignment.CenterVertically
            ) {
                Text(
                    "Foundation 0.1.0",
                    color = MaterialTheme.colorScheme.onSurfaceVariant,
                    style = MaterialTheme.typography.labelSmall
                )
            }
        }
    }
}

@Composable
private fun EngineOption(
    title: String,
    subtitle: String,
    selected: Boolean,
    onClick: () -> Unit
) {
    val colors = if (selected) {
        CardDefaults.cardColors(containerColor = MaterialTheme.colorScheme.secondaryContainer)
    } else {
        CardDefaults.cardColors(containerColor = MaterialTheme.colorScheme.surfaceContainer)
    }
    Card(
        onClick = onClick,
        modifier = Modifier.fillMaxWidth(),
        shape = RoundedCornerShape(18.dp),
        colors = colors
    ) {
        Column(
            modifier = Modifier.padding(16.dp),
            verticalArrangement = Arrangement.spacedBy(4.dp)
        ) {
            Text(title, fontWeight = FontWeight.SemiBold)
            Text(
                subtitle,
                color = MaterialTheme.colorScheme.onSurfaceVariant,
                style = MaterialTheme.typography.bodySmall
            )
        }
    }
}
