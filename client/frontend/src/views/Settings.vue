<template>
  <div class="settings">
    <el-card>
      <template #header>
        <span>应用设置</span>
      </template>
      
      <el-form :model="settings" :rules="rules" ref="formRef" label-width="120px">
        <el-form-item label="服务器地址" prop="serverAddr">
          <el-input v-model="settings.serverAddr" placeholder="例如: 127.0.0.1:10001" />
        </el-form-item>
        
        <el-form-item label="客户端ID" prop="clientId">
          <el-input v-model="settings.clientId" placeholder="自动生成" />
        </el-form-item>
        
        <el-form-item label="自动启动" prop="autoStart">
          <el-switch v-model="settings.autoStart" />
        </el-form-item>
        
        <el-form-item label="日志级别" prop="logLevel">
          <el-select v-model="settings.logLevel" placeholder="请选择日志级别">
            <el-option label="调试" value="debug" />
            <el-option label="信息" value="info" />
            <el-option label="警告" value="warn" />
            <el-option label="错误" value="error" />
          </el-select>
        </el-form-item>
        
        <el-form-item label="连接超时" prop="timeout">
          <el-input-number 
            v-model="settings.timeout" 
            :min="5" 
            :max="300" 
            controls-position="right"
          />
          <span style="margin-left: 10px;">秒</span>
        </el-form-item>
        
        <el-form-item>
          <el-button type="primary" @click="saveSettings">保存设置</el-button>
          <el-button @click="resetSettings">重置</el-button>
        </el-form-item>
      </el-form>
    </el-card>
    
    <el-card style="margin-top: 20px;">
      <template #header>
        <span>关于</span>
      </template>
      
      <div class="about-content">
        <div class="app-info">
          <h3>HiTryRemote 客户端</h3>
          <p>版本: {{ settings.version }}</p>
          <p>基于 QUIC 协议的高性能代理客户端</p>
        </div>
        
        <div class="actions">
          <el-button type="primary" @click="checkUpdate">检查更新</el-button>
          <el-button @click="openLogFolder">打开日志文件夹</el-button>
        </div>
      </div>
    </el-card>
  </div>
</template>

<script setup>
import { ref, reactive, onMounted } from 'vue'
import { ElMessage } from 'element-plus'

const formRef = ref()

const settings = reactive({
  serverAddr: '127.0.0.1:10001',
  clientId: '',
  autoStart: false,
  logLevel: 'info',
  timeout: 30,
  version: '1.0.0'
})

const rules = {
  serverAddr: [
    { required: true, message: '请输入服务器地址', trigger: 'blur' }
  ],
  clientId: [
    { required: true, message: '请输入客户端ID', trigger: 'blur' }
  ],
  logLevel: [
    { required: true, message: '请选择日志级别', trigger: 'change' }
  ]
}

const saveSettings = async () => {
  if (!formRef.value) return
  
  try {
    await formRef.value.validate()
    
    // TODO: 调用后端保存设置
    console.log('保存设置:', settings)
    ElMessage.success('设置保存成功')
  } catch (error) {
    console.error('表单验证失败:', error)
  }
}

const resetSettings = () => {
  // TODO: 重置为默认设置
  settings.serverAddr = '127.0.0.1:10001'
  settings.clientId = 'client-' + Date.now()
  settings.autoStart = false
  settings.logLevel = 'info'
  settings.timeout = 30
  ElMessage.info('设置已重置')
}

const checkUpdate = () => {
  // TODO: 检查更新
  ElMessage.info('当前已是最新版本')
}

const openLogFolder = () => {
  // TODO: 打开日志文件夹
  ElMessage.info('日志文件夹已打开')
}

onMounted(() => {
  // 加载设置
  loadSettings()
})

const loadSettings = () => {
  // TODO: 从后端加载设置
  if (!settings.clientId) {
    settings.clientId = 'client-' + Date.now()
  }
}
</script>

<style scoped>
.settings {
  padding: 20px;
}

.about-content {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.app-info h3 {
  margin: 0 0 10px 0;
  color: #409eff;
}

.app-info p {
  margin: 5px 0;
  color: #666;
}

.actions {
  display: flex;
  gap: 10px;
}
</style>
