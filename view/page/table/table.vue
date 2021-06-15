<template>
  <page-header-wrapper>
    <a-card :bordered="false">
      <div class="table-page-search-wrapper">
        <a-form layout="inline">
          <a-row :gutter="48">
  %s
          </a-row>
        </a-form>
      </div>

      <div class="table-operator">
%s
      </div>

      <s-table
              ref="table"
              size="default"
              rowKey="key"
              :columns="columns"
              :data="loadData"
              :alert="true"
              :rowSelection="rowSelection"
              showPagination="auto"
      >
        <span slot="serial" slot-scope="text, record, index">
          {{ index + 1 }}
        </span>
        <span slot="action" slot-scope="text, record">
          <template>
%s
          </template>
        </span>
      </s-table>

    </a-card>
  </page-header-wrapper>
</template>

<script>
import moment from 'moment'
import { STable } from '@/components'
%s

const columns = [
%s
]

export default {
  name: 'TableList',
  components: {
    STable,
  },
  data () {
    this.columns = columns
    return {
              // create model
              confirmLoading: false,
              // 高级搜索 展开/关闭
              advanced: false,
              // 查询参数
              queryParam: {},
              // 加载数据方法 必须为 Promise 对象
              loadData: parameter => {
                const requestParameters = Object.assign({}, parameter, this.queryParam)
                console.log('loadData request parameters:', requestParameters)
                return getList(requestParameters)
                        .then(res => {
                          console.log(res)
                          return res.result
                        })
              },
%s
            selectedRowKeys: [],
            selectedRows: []
    }
  },

  computed: {
    rowSelection () {
      return {
        selectedRowKeys: this.selectedRowKeys,
        onChange: this.onSelectChange
      }
    }
  },
  methods: {
    handleDelete (id) {
      const _this = this
      this.$confirm({
        title: '你确定删除这个选项吗?',
        content: '删除后将不能恢复，我们将记录您的炒作行为！',
        okText: 'Yes',
        okType: 'danger',
        cancelText: 'No',
        onOk () {
          del(id).then(res => {
            _this.getList()
          })
        }
      })
    },
    onSelectChange (selectedRowKeys, selectedRows) {
      this.selectedRowKeys = selectedRowKeys
      this.selectedRows = selectedRows
    },
    toggleAdvanced () {
      this.advanced = !this.advanced
    },
    resetSearchForm () {
      this.queryParam = {
        date: moment(new Date())
      }
    }
  }
}
</script>
