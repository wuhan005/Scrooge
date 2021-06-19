<template>
  <v-container class="mt-8">
    <v-snackbar v-model="messageBar" color="error" :timeout="2000" :top="true">{{ message }}</v-snackbar>

    <v-row class="text-center">
      <v-col cols="8">
        <Profile/>
        <SponsorList ref="sponsorList"/>
      </v-col>
      <v-col>
        <v-card elevation="1" class="text-left">
          <div v-for="(t, index) in tiers" v-bind:key="index">
            <v-card-title>
              <v-list-item>
                <v-list-item-content>{{ t.amount ? `￥${t.amount}` : '自选金额' }}</v-list-item-content>
                <v-btn depressed color="primary" @click="openDialog(t)">发电</v-btn>
              </v-list-item>
              <v-list-item>
                <div class="body-1">{{ t.comment }}</div>
              </v-list-item>
            </v-card-title>
            <v-divider class="mx-4"></v-divider>
          </div>
        </v-card>
      </v-col>
    </v-row>

    <!-- Payment dialog -->
    <v-dialog v-model="paymentDialog" width="500">
      <v-card>
        <v-stepper v-model="step">
          <v-stepper-items>
            <!-- Form -->
            <v-stepper-content step="1">
              <v-card-title><h3>{{ paymentAmount !== 0 ? `支持 ￥${paymentAmount}` : '自选金额支持' }}</h3></v-card-title>
              <v-card-subtitle>感谢您的支持！</v-card-subtitle>
              <v-card-text>
                <v-text-field v-model="invoiceForm.priceCents" v-show="paymentAmount === 0" placeholder="00.00" dense
                              outlined label="金额"
                              prefix="￥"></v-text-field>
                <v-text-field v-model="invoiceForm.sponsorName" dense outlined label="您的姓名 / 昵称"></v-text-field>
                <v-textarea v-model="invoiceForm.comment" dense outlined label="说点什么？" rows="3"></v-textarea>
              </v-card-text>
              <v-divider></v-divider>
              <v-card-actions>
                <v-spacer></v-spacer>
                <v-btn text color="success" @click="createNewInvoice">
                  <v-icon light>fab fa-weixin</v-icon>
                  <span>微信扫码支付</span>
                </v-btn>
              </v-card-actions>
            </v-stepper-content>
            <!-- Wechat QR Code -->
            <v-stepper-content step="2">
              <v-card-title>
                <h3>微信扫码支付</h3>
              </v-card-title>
              <v-card-subtitle>打开 「微信」 - 「扫一扫」 ，扫描二维码以付款。</v-card-subtitle>

              <v-card-text class="mt-5 mb-5">
                <v-row justify="center">
                  <vue-qr :text="redirectURL"></vue-qr>
                </v-row>
              </v-card-text>
              <v-divider></v-divider>
              <v-card-actions>
                <v-spacer></v-spacer>
                <v-btn color="primary" text @click="closeInvoice">
                  <span>取消支付</span>
                </v-btn>
              </v-card-actions>
            </v-stepper-content>
            <v-stepper-content step="3">
              <v-card-title>
                <h3>支付成功</h3>
              </v-card-title>
              <v-card-text class="text-center">
                <v-icon size="64" color="success">fas fa-check-circle</v-icon>
              </v-card-text>
              <v-card-text class="text-center">
                <span class="body-2 mt-5">感谢您的赞助</span>
              </v-card-text>
            </v-stepper-content>
          </v-stepper-items>
        </v-stepper>

      </v-card>
    </v-dialog>

    <div class="text-center font-weight-light mt-12 body-2">
      <span>Powered by Scrooge / Made with  <v-icon dense color="pink">mdi-heart</v-icon> by E99p1ant</span>
    </div>
  </v-container>
</template>

<script>
import Profile from '@/components/Profile.vue'
import SponsorList from '@/components/SponsorList.vue'

export default {
  name: 'Home',

  data: () => ({
    timer: null,
    messageBar: false,
    message: '',

    paymentDialog: false,
    paymentAmount: 0,
    step: 1,

    tiers: [],

    invoiceForm: {
      priceCents: 0,
      sponsorName: '',
      comment: '',
    },

    redirectURL: '',
    invoiceUID: '',
  }),

  mounted() {
    this.getTiers()
  },

  beforeDestroy() {
    clearInterval(this.timer)
  },

  methods: {
    getTiers() {
      this.utils.GET('/tiers').then(res => {
        this.tiers = res
      }).catch(err => {
        this.messageBar = true
        this.message = err.response.data.msg
      })
    },

    openDialog(t) {
      this.paymentDialog = true;
      this.paymentAmount = t.amount;
      this.invoiceForm = {
        priceCents: t.amount,
        sponsorName: '',
        comment: '',
      }
      this.step = 1;
    },

    createNewInvoice() {
      let priceCents = parseFloat(this.invoiceForm.priceCents)
      if (isNaN(priceCents)) {
        this.messageBar = true
        this.message = '金额不正确'
        return
      }

      this.utils.POST('/pay', {
        PriceCents: priceCents.toFixed(2) * 100,
        SponsorName: this.invoiceForm.sponsorName,
        Comment: this.invoiceForm.comment,
      }).then(res => {
        this.redirectURL = res.redirect_url;
        this.invoiceUID = res.uid;
        this.step = 2;

        this.timer = setInterval(this.checkInvoice, 2000);
      })
    },

    closeInvoice() {
      this.step = 1;
    },

    checkInvoice() {
      this.utils.GET(`/pay/query?uid=${this.invoiceUID}`).then(res => {
        if (res.Paid === true) {
          this.step = 3;
          clearInterval(this.timer)
          this.$refs.sponsorList.$mount();
        }
      })
    }
  },

  components: {
    Profile,
    SponsorList
  },
}
</script>
