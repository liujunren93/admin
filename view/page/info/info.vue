<template>
  <!-- hidden PageHeaderWrapper title demo -->
  <page-header-wrapper :title="false" :content="$t('form.basic-form.basic.description')">
    <a-card :body-style="{padding: '24px 32px'}" :bordered="false">
      <a-form @submit="handleSubmit" :form="form">
        %s
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
%s
export default {
  name: 'BaseForm',
  components: { %s },
  data () {
    return {
      pk: this.$route.query.id,
      %s
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
