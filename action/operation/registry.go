package operation

import (
	"fmt"
	"path"

	"github.com/project-flogo/core/support"
	"github.com/project-flogo/core/support/log"
)

var (
	operations         = make(map[string]Operation)
	operationFactories = make(map[string]Factory)
	operationLoggers   = make(map[string]log.Logger)
)

var operationLogger = log.ChildLogger(log.RootLogger(), "operation")

// Register function registers Operation and it's Factory function.
func Register(operation Operation, f ...Factory) error {

	if operation == nil {
		return fmt.Errorf("cannot register 'nil' operation")
	}

	ref := GetRef(operation)
	ref = path.Base(ref)

	if _, dup := operations[ref]; dup {
		return fmt.Errorf("operation already registered: %s", ref)
	}

	log.RootLogger().Debugf("Registering operation: %s", ref)

	operations[ref] = operation
	name := path.Base(ref) //todo should we use this or the alias?
	operationLoggers[ref] = log.ChildLogger(operationLogger, name)

	if len(f) > 1 {
		log.RootLogger().Warnf("Only one factory can be associated with operation: %s", ref)
	}

	if len(f) == 1 {
		operationFactories[ref] = f[0]
	}

	return nil
}

func GetRef(operation Operation) string {
	return support.GetRef(operation)
}

// Get gets specified activity by ref
func Get(ref string) Operation {
	return operations[ref]
}

// GetFactory gets activity factory by ref
func GetFactory(ref string) Factory {
	return operationFactories[ref]
}

// GetLogger gets activity logger by ref
func GetLogger(ref string) log.Logger {
	if ref[0] == '#' {
		ref, _ = support.GetAliasRef("operation", ref[1:])
	}

	logger, ok := operationLoggers[ref]
	if ok {
		return logger
	} else {
		return log.RootLogger()
	}
}
