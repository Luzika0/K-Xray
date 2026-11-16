package com.panha.kxray

import android.os.Bundle
import androidx.activity.ComponentActivity
import androidx.activity.compose.setContent
import com.panha.kxray.ui.KXrayApp
import com.panha.kxray.ui.theme.KXrayTheme

class MainActivity : ComponentActivity() {
    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
        setContent {
            KXrayTheme {
                KXrayApp()
            }
        }
    }
}
