import { createApp } from "vue"
// @ts-ignore
import App from "@/App.vue"
import PrimeVue from "primevue/config"

import "primevue/resources/themes/lara-light-purple/theme.css"
import "primevue/resources/primevue.min.css"
import "primeicons/primeicons.css"
import "normalize.css"


const app = createApp(App)

app.use(PrimeVue).mount("#app")