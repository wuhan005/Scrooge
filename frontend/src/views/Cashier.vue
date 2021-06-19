<template>
  <v-container v-if="status !== 'pending'">
    <v-card>
      <div v-if="status === 'failed'">
        <v-card-text class="text-center">
          <v-icon size="64" color="pink">fas fa-exclamation-triangle</v-icon>
        </v-card-text>
        <v-card-text class="text-center">
          <span class="body-2 mt-5">支付失败，请点击下方「重新支付」按钮或尝试重新扫描二维码。</span>
        </v-card-text>
        <v-card-text class="mb-5">
          <div class="text-center">
            <v-btn color="primary" @click="cashier">重新支付</v-btn>
          </div>
        </v-card-text>
      </div>
      <div v-if="status === 'cancel'">
        <v-card-text class="text-center">
          <v-icon size="64" color="orange">fas fa-ban</v-icon>
        </v-card-text>
        <v-card-text class="text-center">
          <span class="body-2 mt-5">支付已取消</span>
        </v-card-text>
        <v-card-text class="mb-5">
          <div class="text-center">
            <v-btn color="primary" @click="cashier">重新支付</v-btn>
          </div>
        </v-card-text>
      </div>
      <div v-if="status === 'success'">
        <v-card-text class="text-center">
          <v-icon size="64" color="success">fas fa-check-circle</v-icon>
        </v-card-text>
        <v-card-text class="text-center">
          <h1>￥{{ priceCents / 100 }}</h1>
        </v-card-text>
        <v-card-text class="text-center">
          <span class="body-2 mt-5">支付成功，感谢您的赞助！</span>
        </v-card-text>
      </div>
    </v-card>
    <div class="text-center font-weight-light mt-3 body-2">
      <span>Powered by Scrooge / Made with  <v-icon dense color="pink">mdi-heart</v-icon> by E99p1ant</span>
    </div>

  </v-container>
</template>

<script>
export default {
  name: "Cashier",
  data: () => ({
    status: 'pending',

    uid: '',
    openID: '',

    priceCents: '',
    payment: {
      orderID: '',
      query: '',
    }
  }),

  mounted() {
    this.uid = this.$route.query.uid
    this.openID = this.$route.query.openid

    this.cashier()
  },

  methods: {
    cashier() {
      this.utils.GET(`/pay/cashier?uid=${this.uid}&openID=${this.openID}`).then(res => {
        this.payment.orderID = res.order_id
        this.payment.query = res.query
        this.priceCents = res.price_cents
      }).then(res => {
        this.callWeChatPay()
      })
    },

    callWeChatPay() {
      var vm = this;
      if (typeof WeixinJSBridge == "undefined") {
        if (document.addEventListener) {
          document.addEventListener('WeixinJSBridgeReady', vm.onBridgeReady(vm.payment.query), false);
        } else if (document.attachEvent) {
          document.attachEvent('WeixinJSBridgeReady', vm.onBridgeReady(vm.payment.query));
          document.attachEvent('onWeixinJSBridgeReady', vm.onBridgeReady(vm.payment.query));
        }
      } else {
        vm.onBridgeReady(vm.payment.query);
      }
    },

    onBridgeReady(data) {
      var vm = this;
      WeixinJSBridge.invoke(
          'getBrandWCPayRequest', {
            'appId': data.app_id,
            'timeStamp': data.time_stamp,
            'nonceStr': data.nonce_str,
            'package': data.package,
            'signType': data.sign_type,
            'paySign': data.pay_sign,
          },
          function (res) {
            if (res.err_msg === 'get_brand_wcpay_request:ok') {
              vm.status = 'success';
            } else if (res.err_msg == 'get_brand_wcpay_request:cancel') {
              vm.status = 'cancel';
            } else if (res.err_msg == 'get_brand_wcpay_request:fail') {
              vm.status = 'failed';
            }
          },
      )
    }
  }
}
</script>

<style scoped>

</style>