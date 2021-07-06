<template>
  <page-header-wrapper>
    <a-card :bordered="false">
      <div class="table-page-search-wrapper">
        <a-form layout="inline">
          <a-row :gutter="48">
  		<a-col :md="8" :sm="24">
			<a-form-item label="管理员"> 
 				<a-input v-model="Name" placeholder='请输入管理员'/> 
 			</a-form-item>
		</a-col>
		<a-col :md="!advanced && 8 || 24" :sm="24">
              <span class="table-page-search-submitButtons" :style="advanced && { float: 'right', overflow: 'hidden' } || {} ">
                <a-button type="primary" @click="$refs.table.refresh(true)">查询</a-button>
                <a-button style="margin-left: 8px" @click="() => this.queryParam = {}">重置</a-button>
                <a @click="toggleAdvanced" style="margin-left: 8px">
                  {{ advanced ? '收起' : '展开' }}
                  <a-icon :type="advanced ? 'up' : 'down'"/>
                </a>
              </span>
            </a-col>
          </a-row>
        </a-form>
      </div>

      <div class="table-operator">
		<a-button type="primary" icon="plus"  v-if="$auth('Banner.add')" @click="handleAdd">新建</a-button>
        <a-dropdown v-action:edit v-if="selectedRowKeys.length > 0">
          <a-menu slot="overlay">
            <a-menu-item key="1"  v-if="$auth('Banner.delBatch')"><a-icon type="delete" />删除</a-menu-item>
          </a-menu>
          <a-button style="margin-left: 8px">
            批量操作 <a-icon type="down" />
          </a-button>
        </a-dropdown>
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
		<a  v-if="$auth('Banner.edit')" @click="handleEdit(record)">编辑</a>
		<a-divider  v-if="$auth('Banner.edit')" type="vertical" />
		<a v-if="$auth('Banner.del')" @click="handleDelete(record.id)">删除</a>
          </template>
        </span>
      </s-table>

    </a-card>
  </page-header-wrapper>
</template>

<script>
import moment from 'moment'
import { STable } from '@/components'
import { list } from '@/api/Banner'

const columns = [
	{
		title:'管理员',
		dataIndex:'name',
		key:'name',
	},
	{
		title:'',
		dataIndex:'password',
		key:'password',
	}
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
