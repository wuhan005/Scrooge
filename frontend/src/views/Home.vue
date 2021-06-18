<template>
  <v-container class="mt-8">
    <v-row class="text-center">
      <v-col cols="8">
        <Profile/>
        <SponsorList/>
      </v-col>
      <v-col>
        <v-card elevation="1" class="text-left">
          <div v-for="(t, index) in tiers" v-bind:key="index">
            <v-card-title>
              <v-list-item>
                <v-list-item-content>{{ t.amount ? `￥${t.amount}` : '自选金额' }}</v-list-item-content>
                <v-btn depressed color="primary"
                       @click="paymentDialog = true; paymentAmount = t.amount; step = 1;">发电
                </v-btn>
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
              <v-card-title><h3>{{ paymentAmount ? `支持 ￥${paymentAmount}` : '自选金额支持' }}</h3></v-card-title>
              <v-card-subtitle>感谢您的支持！</v-card-subtitle>
              <v-card-text>
                <v-text-field v-if="paymentAmount === null" placeholder="00.00" dense outlined label="金额"
                              prefix="￥"></v-text-field>
                <v-text-field dense outlined label="您的姓名 / 昵称"></v-text-field>
                <v-textarea dense outlined label="说点什么？" rows="3"></v-textarea>
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
                  <v-img :src="`https://picsum.photos/500/300`" aspect-ratio="1"
                         class="grey lighten-2" max-width="200">
                    <template v-slot:placeholder>
                      <v-row class="fill-height ma-0" align="center" justify="center">
                        <v-progress-circular indeterminate color="grey lighten-5"></v-progress-circular>
                      </v-row>
                    </template>
                  </v-img>
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

  </v-container>
</template>

<script>
import Profile from '@/components/Profile.vue'
import SponsorList from '@/components/SponsorList.vue'

export default {
  name: 'Home',

  data: () => ({
    paymentDialog: false,
    paymentAmount: 0,
    step: 1,

    tiers: [
      {amount: 5, comment: '谢谢老板'},
      {amount: 10, comment: '谢谢老板'},
      {amount: 50, comment: '谢谢老板'},
      {amount: 100, comment: '谢谢老板'},
      {amount: null, comment: '谢谢老板'},
    ],
  }),

  methods: {
    createNewInvoice() {
      this.step = 2;
    },

    closeInvoice() {
      this.step = 1;
    }
  },

  components: {
    Profile,
    SponsorList
  },
}
</script>
