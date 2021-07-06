<template>
  <!-- hidden PageHeaderWrapper title demo -->
  <page-header-wrapper :title="false" :content="$t('form.basic-form.basic.description')">
    <a-card :body-style="{padding: '24px 32px'}" :bordered="false">
      <a-form @submit="handleSubmit" :form="form">
        
		<a-fo1rm-item
          :label="管理员"
          :labelCol="{lg: {span: 7}, sm: {span: 7}}"
          :wrapperCol="{lg: {span: 10}, sm: {span: 17} }"> 
			<a-select v-decorator="['name', {initialValue:'' }]" :options="bannerData" placeholder="请选择管理员" />
		</a-form-item>
		<a-fo1rm-item
          :label="密码"
          :labelCol="{lg: {span: 7}, sm: {span: 7}}"
          :wrapperCol="{lg: {span: 10}, sm: {span: 17} }"> 
			<a-select v-decorator="['password', {initialValue:'value' }]" :options="[{"label":"label","value":"value"},{"label":"label1","value":"value1"}]Data" placeholder="请选择密码" />
		</a-form-item>
        <a-form-item
          :wrapperCol="{ span: 24 }"
          style="text-align: center"
        >
          <a-button htmlType="submit" type="primary">{{ $t('form.basic-form.form.submit') }}</a-button>
          <a-button style="margin-left: 8px">{{ $t('form.basic-form.form.save') }}</a-button>
        </a-form-item>
      </a-form>
    </a-card>

  </page-header-wrapper>
</template>

<script>
import { list as bannerList } from '@/api/Banner'
import { create,info,update } from '@/api/Admin'
export default {
  name: 'BaseForm',
  components: {  },
  data () {
    return {
      pk: this.$route.query.id,
      	bannerData: () => { return  bannerList().then(res => {return res.data})},
      form: this.$form.createForm(this)
    }
  },
  methods: {
    info () {
      if (!this.pk) {
          return false
      }
      info(this.pk).then(res => {
          this.form.setFieldsValue(res.data)
      })
    },
    // handler
    handleSubmit (e) {
      e.preventDefault()
      this.form.validateFields((err, values) => {
        console.log(values)
        if (err) {
          // console.log('Received values of form: ', values)
          return false
        }
           if (this.pk) {
            update(this.$route.query.id, values)
        } else {
            create(values)
        }
      })
    }
  }
}
</script>
