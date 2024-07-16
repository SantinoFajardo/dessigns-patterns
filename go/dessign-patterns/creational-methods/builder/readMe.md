<div>
<h1>Builder:</h1>
<h2>Definition:</h2>
<p>The Builder is a creational design pattern that provides an interface for constructing complex objects step by step. Unlike other creational patterns, which construct products in a single step, the Builder pattern focuses on constructing a complex product incrementally. It allows for more control over the construction process and can produce different representations of the object using the same construction process.</p>
<h2>Implementation:</h2>
<p>In the Builder pattern, we create an interface that declares a set of methods for constructing parts of a complex object. Concrete builders implement this interface to build and assemble the parts of the product. A director class defines the order in which to call the construction steps, and the client uses the director to get the final product.</p>
<h2>How it works:</h2>
<p>The Builder works by defining an abstract builder interface with methods for creating parts of the product. Concrete builder classes implement these methods to construct and assemble the specific parts. A director class oversees the construction process, ensuring that the product is built in a specific order. This approach separates the construction of the product from its representation, allowing the same construction process to create different representations of the product.</p>
<h2>What can we solve with the builder:</h2>
<h3>A real world problem:</h3>
<p>Imagine a restaurant that offers customized meals. Each meal can have different components, such as a main course, side dish, drink, and dessert. The restaurant needs a way to create different types of meals without tightly coupling the client code to the specific meal classes.</p>
<h3>Solution implemented with the builder:</h3>
<p>The Builder can be used to solve this problem by creating an interface `MealBuilder` with methods `addMainCourse()`, `addSideDish()`, `addDrink()`, and `addDessert()`. Concrete builders such as `VegetarianMealBuilder`, `NonVegetarianMealBuilder`, and `VeganMealBuilder` would implement this interface to create instances of specific meals. A `MealDirector` class would define the order in which to call the methods of the builder to construct different types of meals. The client code would use the `MealDirector` and `MealBuilder` interface to create customized meals without knowing the specific classes behind them, thus promoting flexibility and scalability.</p>
<h2>Some pros and cons</h2>
<h3>Pros:</h3>
<ul>
<li>Allows for step-by-step construction of complex objects, providing more control over the construction process.</li>
<li>Separates the construction of the product from its representation, enabling different representations to be built using the same construction process.</li>
<li>Encapsulates the construction logic, making the code more maintainable and easier to understand.</li>
<li>Adheres to the Single Responsibility Principle by allowing the construction process to be separated from the product's representation.</li>
</ul>
<h3>Cons:</h3>
<ul>
<li>Can increase the complexity of the code with the addition of multiple builder classes and products.</li>
<li>May lead to a proliferation of classes, which can be harder to manage and understand.</li>
<li>Sometimes overkill for simple object construction tasks where a simple factory method might suffice.</li>
</ul>
<h2>When to use it and when not:</h2>
<h3>When to use it:</h3>
<ul>
<li>When you need to construct complex objects that require multiple steps or configurations.</li>
<li>When you want to separate the construction process from the representation of the product.</li>
<li>When you need to create different representations of the same product using the same construction process.</li>
</ul>
<h3>When not to use it:</h3>
<ul>
<li>When the objects being created do not require multiple steps or configurations.</li>
<li>When the object construction is straightforward and does not involve any complex logic or dependencies.</li>
<li>When the overhead of creating and maintaining multiple builder classes outweighs the benefits of control and separation of construction logic.</li>
</ul>
<h4>Graphic representation: </h4>
<img src="https://refactoring.guru/images/patterns/diagrams/builder/structure.png"/>
</div>
