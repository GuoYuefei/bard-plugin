/**
	此包是插件包，包含一个控制子协议的子包（控制子协议也是一种插件）
	插件分成两种
	一种是对流量进行操控的
	一种是对本次发包进行控制的。。。。现在想到的就是本次write的字节数是要被接收方知道
	原因在于发送是对一组数据操作，接受是也应该还原到原本的数据然后进行反操作
 */
package bard_plugin