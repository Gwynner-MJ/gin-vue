<template>
  <div class="regiser">
    <b-row class="mt-5">
      <b-col
        md="8"
        offset-md="2"
        lg="6"
        offset-lg="3"
      >
        <b-card title="登录">
          <b-form>
            <b-form-group label="手机号">
              <b-form-input
                v-model="$v.user.telephone.$model"
                type="tel"
                placeholder="输入电话号码"
                :state="validatestate('telephone')"
              ></b-form-input>
              <b-form-invalid-feedback :state="validatestate('telephone')">
                手机号不符合规范
              </b-form-invalid-feedback>
            </b-form-group>
            <b-form-group label="密码">
              <b-form-input
                v-model="$v.user.password.$model"
                type="password"
                placeholder="输入用户密码"
                :state="validatestate('password')"
              ></b-form-input>
              <b-form-invalid-feedback :state="validatestate('password')">
                密码必须≥6位
              </b-form-invalid-feedback>
            </b-form-group>
            <b-button
              @click="login"
              variant="outline-primary"
              block
            >登录</b-button>
            <b-form-group>
            </b-form-group>
          </b-form>
        </b-card>
      </b-col>
    </b-row>
  </div>
</template>
<script>
import { required, minLength } from 'vuelidate/lib/validators';
import customvalidator from '@/helper/validator';

export default {
  data() {
    return {
      user: {
        telephone: '',
        password: '',
      },
    };
  },
  validations: {
    user: {
      telephone: {
        required,
        telephone: customvalidator.telephonevalidator,
      },
      password: {
        required,
        minLength: minLength(6),
      },
    },
  },
  methods: {
    validatestate(name) {
      const { $dirty, $error } = this.$v.user[name];
      return $dirty ? !$error : null;
    },
    login() {
      console.log('login');
    },
  },
};
</script>
<style lang="scss" scoped>
</style>
