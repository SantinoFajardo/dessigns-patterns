<div>
<h1>Abstract Factory:</h1>
<h2>Definition:</h2>
<p>The Abstract Factory is a creational design pattern that provides an interface for creating families of related or dependent objects without specifying their concrete classes. It is used to create sets of related objects, or families, by implementing multiple factory methods in a single factory interface. This pattern promotes consistency among objects created by ensuring that they are all from the same family.</p>
<h2>Implementation:</h2>
<p>In the Abstract Factory pattern, we create an interface that declares a set of methods for creating abstract products. Concrete factories implement this interface to produce specific products. Each concrete factory corresponds to a specific family of products. The client code uses the abstract factory to create product objects, ensuring that they belong to the same family.</p>
<h2>How it works:</h2>
<p>The Abstract Factory works by defining an abstract factory interface with methods for creating abstract products. Concrete factory classes implement these methods to instantiate specific products. This ensures that the client code uses a common interface to create related objects, promoting consistency and reducing dependency on specific classes.</p>
<h2>What can we solve with the abstract factory:</h2>
<h3>A real world problem:</h3>
<p>Imagine a furniture store that sells different styles of furniture, such as modern, Victorian, and art deco. Each style has its own set of products, including chairs, tables, and sofas. The furniture store needs a way to create instances of these products without tightly coupling the client code to the specific product classes.</p>
<h3>Solution implemented with the abstract factory:</h3>
<p>The Abstract Factory can be used to solve this problem by creating an interface `FurnitureFactory` with methods `createChair()`, `createTable()`, and `createSofa()`. Subclasses such as `ModernFurnitureFactory`, `VictorianFurnitureFactory`, and `ArtDecoFurnitureFactory` would implement this interface to create instances of ModernChair, ModernTable, ModernSofa, and so on. The client code would use the `FurnitureFactory` interface to create the required furniture pieces without knowing the specific classes behind them, thus promoting flexibility and scalability.</p>
<h2>Some pros and cons</h2>
<h3>Pros:</h3>
<ul>
<li>Promotes consistency among related products by ensuring they are created from the same family.</li>
<li>Encapsulates the creation of related products, reducing the dependency on specific implementations.</li>
<li>Adheres to the Open/Closed Principle by allowing new product families to be added without modifying existing code.</li>
</ul>
<h3>Cons:</h3>
<ul>
<li>Can increase the complexity of the code with the addition of multiple factory classes and products.</li>
<li>May lead to a proliferation of classes, which can be harder to manage and understand.</li>
<li>Sometimes overkill for simple object creation tasks where a simple factory method might suffice.</li>
</ul>
<h2>When to use it and when not:</h2>
<h3>When to use it:</h3>
<ul>
<li>When you need to create families of related objects that must be used together.</li>
<li>When you want to ensure consistency among products created by a factory.</li>
<li>When the creation process involves multiple steps or configurations that are common to all products in a family.</li>
</ul>
<h3>When not to use it:</h3>
<ul>
<li>When the objects being created do not belong to a family or do not need to be consistent with each other.</li>
<li>When the object creation is straightforward and does not involve any complex logic or dependencies.</li>
<li>When the overhead of creating and maintaining multiple factory classes outweighs the benefits of consistency and decoupling.</li>
</ul>
<h4>Grafic representation: </h4>
<img src="./assets/Screenshot 2024-07-04 at 10.02.15.png"/>
</div>
