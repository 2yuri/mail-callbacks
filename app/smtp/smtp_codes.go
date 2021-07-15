package smtp

type smtpCodes struct {
	Code        string
	Description string
}

type Status struct{
	codes map[string]smtpCodes
}

var smtpStatus *Status

func GetSTMPStatus() *Status {
	return smtpStatus
}

func (s *Status) Codes() map[string]smtpCodes {
	return s.codes
}

func InitSMTPCodes() {

	codes := make(map[string]smtpCodes)

	codes["X.0.0"] = smtpCodes{
		Code:        "Status indefinido",
		Description: "Código de erro desconhecido. Não foi possível identificar a causa do erro.",
	}
	codes["X.1.0"] = smtpCodes{
		Code:        "Problemas com o endereço",
		Description: "Alguma informação no endereço informado pode conter erro, verifique.",
	}
	codes["X.1.1"] = smtpCodes{
		Code:        "Caixa de email destino errada",
		Description: "A caixa de e-mail informada não existe, verifique a porção do endereço do lado esquerdo do @.",
	}
	codes["X.1.10"] = smtpCodes{
		Code:        "Caixa de email destino errada",
		Description: "A caixa de e-mail informada não existe, verifique a porção do endereço do lado esquerdo do @.",
	}
	codes["X.1.2"] = smtpCodes{
		Code:        "Problema no dominio de email destino",
		Description: "O endereço de dominio informado não existe ou está incapaz de aceitar mensagens. Verifique a porção de endereço do lado direto do @.",
	}
	codes["X.1.3"] = smtpCodes{
		Code:        "Sintaxe de endereço de caixa de correio de destino inválida",
		Description: "Verifique o endereço de e-mail, existem informações com sintaxe inválida.",
	}
	codes["X.1.4"] = smtpCodes{
		Code:        "Endereço da caixa de correio de destino ambíguo",
		Description: "O endereço da caixa de correio conforme especificado corresponde a um ou mais destinatários no sistema de destino.",
	}
	codes["X.1.5"] = smtpCodes{
		Code:        "Endereço de destino válido",
		Description: "Este endereço de caixa de correio, conforme especificado, era válido.",
	}
	codes["X.1.6"] = smtpCodes{
		Code:        "Endereço de e-mail foi movido, sem endereço de encaminhamento",
		Description: "O endereço de caixa de correio fornecido já foi válido, mas o e-mail é não sendo mais aceito para aquele endereço.",
	}
	codes["X.1.7"] = smtpCodes{
		Code:        "Sintaxe de endereço de caixa de correio de remetente incorreto",
		Description: "O endereço do remetente era sintaticamente inválido. Isso pode se aplicar a qualquer campo do endereço.",
	}
	codes["X.1.8"] = smtpCodes{
		Code:        "Endereço do sistema do remetente incorreto",
		Description: "O sistema do remetente especificado no endereço não existe ou é incapaz de aceitar correspondência de retorno. Para nomes de domínio, este significa que a parte do endereço à direita do @ é inválida para o correio.",
	}
	codes["X.2.0"] = smtpCodes{
		Code:        "Outro status de caixa de correio ou indefinido",
		Description: "A caixa de correio existe, mas algo sobre a caixa de correio de destino causou o envio deste erro.",
	}
	codes["X.2.1"] = smtpCodes{
		Code:        "Caixa de correio desativada, não aceita mensagens",
		Description: "A caixa de correio existe, mas não está aceitando mensagens. Isso pode ser um erro permanente se a caixa de correio nunca for reativada ou um erro transitório se a caixa de correio estiver desativada apenas temporariamente.",
	}
	codes["X.2.2"] = smtpCodes{
		Code:        "Caixa de correio cheia",
		Description: "A caixa de correio está cheia porque o usuário excedeu uma cota administrativa ou capacidade física por caixa de correio. A semântica geral implica que o destinatário pode excluir mensagens para disponibilizar mais espaço.",
	}
	codes["X.2.3"] = smtpCodes{
		Code:        "O comprimento da mensagem excede o limite administrativo",
		Description: "Um limite de comprimento de mensagem administrativa por caixa de correio foi ultrapassado.",
	}
	codes["X.2.4"] = smtpCodes{
		Code:        "Problema de expansão da lista de correio",
		Description: "A caixa de correio é um endereço de lista de discussão e a lista de discussão era incapaz de ser expandido.",
	}
	codes["X.3.0"] = smtpCodes{
		Code:        "Outro status do sistema de e-mail ou indefinido",
		Description: "O sistema de destino existe e normalmente aceita correio, mas algo sobre o sistema causou a geração deste erro.",
	}
	codes["X.3.1"] = smtpCodes{
		Code:        "Sistema de correio cheio",
		Description: "O armazenamento do sistema de correio foi excedido. A semântica geral implica que o destinatário individual pode não ser capaz de excluir material para abrir espaço para mensagens adicionais.",
	}
	codes["X.3.2"] = smtpCodes{
		Code:        "Sistema não aceita mensagens de rede",
		Description: "O host no qual a caixa de correio está residente não está aceitando mensagens. Exemplos de tais condições incluem um imanente desligamento, carga excessiva ou manutenção do sistema.",
	}
	codes["X.3.3"] = smtpCodes{
		Code:        "Sistema não capaz de recursos selecionados",
		Description: "Os recursos selecionados especificados para a mensagem não são suportados pelo sistema de destino. Isso pode ocorrer em gateways quando recursos de um domínio não podem ser mapeados para os recurso em outro.",
	}
	codes["X.3.4"] = smtpCodes{
		Code:        "Mensagem muito grande para o sistema",
		Description: "A mensagem é maior do que o limite de tamanho por mensagem. Este limite pode ser por razões físicas ou administrativas.",
	}
	codes["X.3.5"] = smtpCodes{
		Code:        "Sistema configurado incorretamente",
		Description: "O sistema não está configurado de maneira que permita aceitar esta mensagem.",
	}
	codes["X.4.0"] = smtpCodes{
		Code:        "Outra rede ou status de roteamento ou rede indefinida",
		Description: "Algo deu errado com a rede, mas não está claro qual é o problema, ou o problema não pode ser bem expresso com qualquer um dos outros códigos de detalhe fornecidos.",
	}
	codes["X.4.1"] = smtpCodes{
		Code:        "Sem resposta do anfitrião",
		Description: "A tentativa de conexão de saída não foi respondida, porque ou o sistema remoto estava ocupado ou não conseguiu realizar um ligar.",
	}
	codes["X.4.14"] = smtpCodes{
		Code:        "Sem resposta do anfitrião",
		Description: "A tentativa de conexão de saída não foi respondida, porque ou o sistema remoto estava ocupado ou não conseguiu realizar um ligar.",
	}
	codes["X.4.2"] = smtpCodes{
		Code:        "Conexão ruim",
		Description: "A conexão de saída foi estabelecida, mas não foi possível completar a transação da mensagem, seja devido ao tempo limite, ou qualidade de conexão inadequada.",
	}
	codes["X.4.3"] = smtpCodes{
		Code:        "Falha do servidor de diretório",
		Description: "O sistema de rede não foi capaz de encaminhar a mensagem, porque um o servidor de diretório não estava disponível.",
	}
	codes["X.4.4"] = smtpCodes{
		Code:        "Incapaz de rotear",
		Description: "O sistema de correio não foi capaz de determinar o próximo salto para o mensagem porque as informações de roteamento necessárias eram indisponível no servidor de diretório.",
	}
	codes["X.4.5"] = smtpCodes{
		Code:        "Congestionamento do sistema de correio",
		Description: "O sistema de e-mail não conseguiu entregar a mensagem porque o sistema de correio estava congestionado.",
	}
	codes["X.4.6"] = smtpCodes{
		Code:        "Loop de roteamento detectado",
		Description: "Um loop de roteamento fez com que a mensagem fosse encaminhada muitas vezes, seja por causa de tabelas de roteamento incorretas ou um usuário loop de encaminhamento.",
	}
	codes["X.4.7"] = smtpCodes{
		Code:        "Tempo de entrega expirou",
		Description: "A mensagem foi considerada muito antiga pelo sistema de rejeição, seja porque permaneceu no host por muito tempo ou porque o valor de tempo de vida especificado pelo remetente da mensagem foi excedido.",
	}
	codes["X.5.0"] = smtpCodes{
		Code:        "Outro ou status de protocolo indefinido",
		Description: "Algo estava errado com o protocolo necessário para entregar o mensagem para o próximo salto e o problema não pode ser bem expresso com qualquer um dos outros códigos de detalhe fornecidos.",
	}
	codes["X.5.1"] = smtpCodes{
		Code:        "Comando inválido",
		Description: "Foi emitido um comando de protocolo de transação de correio que era fora da sequência ou sem suporte.",
	}
	codes["X.5.2"] = smtpCodes{
		Code:        "Erro de sintaxe",
		Description: "Foi emitido um comando de protocolo de transação de correio que não pode ser interpretado, seja porque a sintaxe estava errada ou porque comando não é reconhecido.",
	}
	codes["X.5.3"] = smtpCodes{
		Code:        "Muitos recipientes",
		Description: "Mais destinatários foram especificados para a mensagem do que poderiam ter sido entregues pelo protocolo. Este erro normalmente deve resultar na segmentação da mensagem em dois, o restante dos destinatários a serem entregues em uma tentativa de entrega subsequente.",
	}
	codes["X.5.4"] = smtpCodes{
		Code:        "Argumentos de comando inválidos",
		Description: "Um comando de protocolo de transação de correio válido foi emitido com argumentos inválidos, seja porque os argumentos estavam fora de intervalo ou representados recursos não reconhecidos.",
	}
	codes["X.5.5"] = smtpCodes{
		Code:        "Versão de protocolo errada",
		Description: "Existia uma versão de protocolo incompatível que não poderia ser resolvido automaticamente pelas partes comunicantes.",
	}
	codes["X.6.0"] = smtpCodes{
		Code:        "Outro ou erro de mídia indefinido",
		Description: "Algo sobre o conteúdo de uma mensagem fez com que ela fosse considerado impossível de ser entregue e o problema não pode estar bem expresso com qualquer um dos outros códigos de detalhe fornecidos.",
	}
	codes["X.6.1"] = smtpCodes{
		Code:        "Mídia não suportada",
		Description: "O conteúdo da mensagem deve ser convertido antes que possa ser entregue e tal conversão não é permitida. Tal proibições podem ser a expressão do remetente na mensagem próprio ou a política do host de envio.",
	}
	codes["X.6.2"] = smtpCodes{
		Code:        "Conversão necessária e proibida",
		Description: "O conteúdo da mensagem deve ser convertido antes que possa ser entregue e tal conversão não é permitida. Tal proibições podem ser a expressão do remetente na mensagem próprio ou a política do host de envio.",
	}
	codes["X.6.3"] = smtpCodes{
		Code:        "Conversão necessária, mas não suportada",
		Description: "O conteúdo da mensagem deve ser convertido para ser encaminhado, mas essa conversão não é possível ou prática por um host no caminho de encaminhamento.",
	}
	codes["X.6.4"] = smtpCodes{
		Code:        "Conversão com perda realizada",
		Description: "Este é um aviso enviado ao remetente quando a entrega da mensagem foi bem-sucedida, mas quando a entrega exigiu uma conversão de midia na qual alguns dados foram perdidos.",
	}
	codes["X.6.5"] = smtpCodes{
		Code:        "Falha na conversão",
		Description: "Uma conversão de midia foi necessária, mas não teve êxito.",
	}
	codes["X.7.0"] = smtpCodes{
		Code:        "Outro status de segurança ou indefinido",
		Description: "Algo relacionado à segurança fez com que a mensagem fosse retornou, e o problema não pode ser bem expresso com qualquer um dos os outros códigos de detalhe fornecidos.",
	}
	codes["X.7.1"] = smtpCodes{
		Code:        "Entrega não autorizada, mensagem recusada",
		Description: "O remetente não está autorizado a enviar para o destino. Isso pode ser o resultado da filtragem por host ou por destinatário.",
	}
	codes["X.7.2"] = smtpCodes{
		Code:        "Expansão da lista de correio proibida",
		Description: "O remetente não está autorizado a enviar uma mensagem para o destinatário lista de discussão.",
	}
	codes["X.7.3"] = smtpCodes{
		Code:        "Conversão de segurança necessária, mas não possível",
		Description: "Uma conversão de um protocolo de mensagens seguras para outro foi necessário para entrega e tal conversão não foi possível.",
	}
	codes["X.7.4"] = smtpCodes{
		Code:        "Recursos de segurança não suportados",
		Description: "Uma mensagem continha recursos de segurança, como autenticação que não pode ser suportada na entrega protocolo.",
	}
	codes["X.7.5"] = smtpCodes{
		Code:        "Falha criptográfica",
		Description: "Um sistema de transporte de outra forma autorizado para validar ou descriptografar uma mensagem no transporte não foi capaz de fazer isso porque necessário informações como chave não estavam disponíveis ou tais informações era inválido.",
	}
	codes["X.7.6"] = smtpCodes{
		Code:        "Algoritmo criptográfico não compatível",
		Description: "Um sistema de transporte de outra forma autorizado para validar ou descriptografar uma mensagem não foi capaz de fazer isso porque o algoritmo necessário não era compatível.",
	}
	codes["X.7.7"] = smtpCodes{
		Code:        "Falha de integridade de mensagem",
		Description: "Um sistema de transporte de outra forma autorizado para validar uma mensagem não foi capaz de fazer isso porque a mensagem estava corrompida ou alterado. Isso pode ser útil como um permanente, temporário persistente ou código de entrega bem-sucedido.",
	}
	codes["X.X.X"] = smtpCodes{
		Code:        "Erro genérico",
		Description: "Erro genérico sem detalhamento.",
	}

	smtpStatus = &Status{
		codes: codes,
	}
}
