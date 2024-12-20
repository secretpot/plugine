package base

// BasePlugin 提供了 Plugin 接口的基础实现
type BasePlugin struct {
	name     string
	settings interface{}
}

// Plugin 定义了插件必须实现的接口
type Plugin interface {
	// 获取插件名称
	GetName() string
	// 获取插件设置
	GetSettings() interface{}
	// 设置插件名称
	SetName(name string)
	// 设置插件配置
	SetSettings(settings interface{})

	// 插件生命周期钩子
	Before() error
	After() error
	Callback() error
}

// 实现 BasePlugin 的基础方法
func (b *BasePlugin) GetName() string {
	return b.name
}

func (b *BasePlugin) GetSettings() interface{} {
	return b.settings
}

func (b *BasePlugin) SetName(name string) {
	b.name = name
}

func (b *BasePlugin) SetSettings(settings interface{}) {
	b.settings = settings
}

// 提供默认的空实现，允许继承者只实现需要的方法
func (b *BasePlugin) Before() error   { return nil }
func (b *BasePlugin) After() error    { return nil }
func (b *BasePlugin) Callback() error { return nil }
