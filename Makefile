puml:
	goplantuml -recursive creational/singleton > creational/singleton/singleton.puml
	goplantuml -recursive creational/factory > creational/factory/factory.puml
	goplantuml -recursive creational/abstract_factory > creational/abstract_factory/abstract_factory.puml
	goplantuml -recursive creational/builder > creational/builder/builder.puml
	goplantuml -recursive creational/prototype > creational/prototype/prototype.puml

	goplantuml -recursive structural/adapter > structural/adapter/adapter.puml
	goplantuml -recursive structural/composition > structural/composition/composition.puml
