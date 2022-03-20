<template>
  <div class="regiser">
    <b-row class="mt-5">
      <b-col
        md="8"
        offset-md="2"
        lg="6"
        offset-lg="3"
      >
        <b-card title="注册">
          <b-form>
            <b-form-group label="姓名">
              <b-form-input
                v-model="$v.user.name.$model"
                type="text"
                placeholder="输入用户名称(选填)"
              ></b-form-input>
            </b-form-group>
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
              @click="register"
              variant="outline-primary"
              block
            >注册</b-button>
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
        name: '',
        telephone: '',
        password: '',
      },
    };
  },
  validations: {
    user: {
      name: {},
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
      // 这里是es6 解构
      const { $dirty, $error } = this.$v.user[name];
      return $dirty ? !$error : null;
    },
    register() {
      // 验证数据
      this.$v.user.$touch();
      if (this.$v.user.$anyError) {
        return;
      }
      // 请求
      const api = 'http://localhost:1060/api/auth/register';
      this.axios.post(api, { ...this.user }).then((res) => {
        // 保存token
        console.log(res.data);
        // 跳转主页
      }).catch((err) => {
        console.log('err:', err.response.data.msg);
      });
      console.log('register');
    },
  },
};
</script>
<style lang="scss" scoped>
</style>
